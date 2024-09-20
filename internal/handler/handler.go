package handler

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"github.com/egor-zakharov/goph-keeper/internal/service/users"
	usersStorage "github.com/egor-zakharov/goph-keeper/internal/storage/users"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServerServer
	usersService users.Service
	cardsService cards.Service
	syncClients  map[string]map[string]pb.GophKeeperServer_SubscribeToChangesServer
	rwMutex      sync.RWMutex
}

func New(service users.Service, cardsService cards.Service) *GophKeeperServer {
	return &GophKeeperServer{
		usersService: service,
		cardsService: cardsService,
		syncClients:  make(map[string]map[string]pb.GophKeeperServer_SubscribeToChangesServer),
	}
}

func (s *GophKeeperServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	response := &pb.SignUpResponse{}

	user := models.User{
		Login:    in.Login,
		Password: in.Password,
	}
	if !user.IsValidLogin() || !user.IsValidPass() {
		logger.Log().Sugar().Errorw("SignUp handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Login or password should not be empty")
	}
	createdUser, err := s.usersService.Register(ctx, user)

	if errors.Is(err, usersStorage.ErrConflict) {
		logger.Log().Sugar().Errorw("SignUp handler", "usersService register", err)
		return response, status.Errorf(codes.InvalidArgument, "User with such login already exists")
	}

	sessionID := uuid.New().String()
	JWTToken, err := auth.BuildJWTString(createdUser.UserID, sessionID)

	if err != nil {
		logger.Log().Sugar().Errorw("SignUp handler", "build jwt", err)
		return response, status.Errorf(codes.Internal, "Can not build auth token")
	}

	response.Token = JWTToken
	return response, nil

}

func (s *GophKeeperServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	response := &pb.SignInResponse{}

	user := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	if !user.IsValidLogin() || !user.IsValidPass() {
		logger.Log().Sugar().Errorw("SignUp handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Login or password should not be empty")
	}

	usr, err := s.usersService.Login(ctx, user)

	if err != nil {
		logger.Log().Sugar().Errorw("SignIn handler", "usersService login", err)
		return response, status.Errorf(codes.InvalidArgument, "Invalid login or password")
	}

	sessionID := uuid.New().String()
	JWTToken, err := auth.BuildJWTString(usr.UserID, sessionID)
	if err != nil {
		logger.Log().Sugar().Errorw("SignIn handler", "build jwt", err)
		return response, status.Errorf(codes.Internal, "Can not build auth token")
	}

	response.Token = JWTToken
	return response, nil
}

func (s *GophKeeperServer) CreateCard(ctx context.Context, in *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	response := &pb.CreateCardResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	if in.Card == nil {
		return nil, nil
	}

	card := models.Card{
		Number:         in.Card.Number,
		ExpirationDate: in.Card.ExpirationDate,
		HolderName:     in.Card.HolderName,
		CVV:            in.Card.Cvv,
	}

	if !card.IsValidNumber() || !card.IsValidDate() || !card.IsValidNumber() {
		logger.Log().Sugar().Errorw("CreateCard handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Incorrect card data")
	}

	createdCard, err := s.cardsService.Create(ctx, card, userID)
	if err != nil {
		return nil, err
	}

	response.CardID = createdCard.ID
	s.sendNotifications(ctx, "card", "create", response.CardID)
	return response, nil
}

func (s *GophKeeperServer) GetCards(ctx context.Context, _ *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	response := &pb.GetCardsResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	cards, err := s.cardsService.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, card := range *cards {
		response.Cards = append(response.Cards, &pb.GetCardsResponse_Card{
			Id:             card.ID,
			Number:         card.Number,
			ExpirationDate: card.ExpirationDate,
			HolderName:     card.HolderName,
			Cvv:            card.CVV,
		})
	}

	return response, nil
}

func (s *GophKeeperServer) UpdateCard(ctx context.Context, in *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	if in.Card == nil {
		return nil, nil
	}

	userID := ctx.Value(auth.UserIdContextKey).(string)
	card := models.Card{
		ID:             in.Card.Id,
		Number:         in.Card.Number,
		ExpirationDate: in.Card.ExpirationDate,
		HolderName:     in.Card.HolderName,
		CVV:            in.Card.Cvv,
	}
	if !card.IsValidNumber() || !card.IsValidDate() || !card.IsValidNumber() {
		logger.Log().Sugar().Errorw("UpdateCard handler", "validation error")
		return &pb.UpdateCardResponse{Result: false}, status.Errorf(codes.InvalidArgument, "Incorrect card data")
	}
	_, err := s.cardsService.Update(ctx, card, userID)
	if err != nil {
		return nil, err
	}
	s.sendNotifications(ctx, "card", "update", in.Card.Id)

	return &pb.UpdateCardResponse{Result: true}, nil
}

func (s *GophKeeperServer) DeleteCard(ctx context.Context, in *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := s.cardsService.Delete(ctx, in.Id, userID)
	if err != nil {
		return &pb.DeleteCardResponse{Result: false}, err
	}
	s.sendNotifications(ctx, "card", "delete", in.Id)
	return &pb.DeleteCardResponse{Result: true}, err
}

// SubscribeToChanges - stream changes to clients
func (s *GophKeeperServer) SubscribeToChanges(in *pb.SubscribeToChangesRequest, stream pb.GophKeeperServer_SubscribeToChangesServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	userID, _ := ctx.Value(auth.UserIdContextKey).(string)
	sessionID, _ := ctx.Value(auth.SessionIdContextKey).(string)
	defer cancel()
	s.rwMutex.Lock()
	if len(s.syncClients[userID]) == 0 {
		s.syncClients[userID] = make(map[string]pb.GophKeeperServer_SubscribeToChangesServer)
	}
	s.syncClients[userID][sessionID] = stream
	s.rwMutex.Unlock()
	for {
		time.Sleep(time.Minute)
	}
}

// SendNotifications - stream all user session with update
func (s *GophKeeperServer) sendNotifications(ctx context.Context, product string, action string, ID string) {
	sessionID, _ := ctx.Value(auth.SessionIdContextKey).(string)
	userID, _ := ctx.Value(auth.UserIdContextKey).(string)
	s.rwMutex.Lock()
	for session, client := range s.syncClients[userID] {
		if session != sessionID {
			_ = client.Send(&pb.SubscribeToChangesResponse{
				Product: product,
				Action:  action,
				Id:      ID,
			})
		}
	}
	s.rwMutex.Unlock()
}
