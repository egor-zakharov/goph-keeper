package server

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signin"
	"github.com/egor-zakharov/goph-keeper/internal/handlers/signup"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	authService "github.com/egor-zakharov/goph-keeper/internal/service/auth"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServerServer
	signUp       *signup.Handler
	signIn       *signin.Handler
	cardsService cards.Service
	authService  authService.Service
	syncClients  map[string]map[string]pb.GophKeeperServer_SubscribeToChangesServer
	rwMutex      sync.RWMutex
}

func New(cardsService cards.Service, signUp *signup.Handler, signIn *signin.Handler, authService authService.Service) *GophKeeperServer {
	return &GophKeeperServer{
		cardsService: cardsService,
		signUp:       signUp,
		signIn:       signIn,
		authService:  authService,
		syncClients:  make(map[string]map[string]pb.GophKeeperServer_SubscribeToChangesServer),
	}
}

func (s *GophKeeperServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return s.signUp.SignUp(ctx, in)
}

func (s *GophKeeperServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return s.signIn.SignIn(ctx, in)
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

func (s *GophKeeperServer) CreateAuthData(ctx context.Context, in *pb.CreateAuthDataRequest) (*pb.CreateAuthDataResponse, error) {
	response := &pb.CreateAuthDataResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)
	if in.Data == nil {
		return nil, nil
	}
	authData := models.AuthData{
		Meta:     in.Data.Meta,
		Login:    in.Data.Login,
		Password: in.Data.Password,
	}
	data, err := s.authService.Create(ctx, authData, userID)
	if err != nil {
		return nil, err
	}
	response.Id = data.ID
	s.sendNotifications(ctx, authData.Meta, "create", response.Id)
	return response, nil
}

func (s *GophKeeperServer) GetAuthData(ctx context.Context, in *pb.GetAuthDataRequest) (*pb.GetAuthDataResponse, error) {
	response := &pb.GetAuthDataResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	data, err := s.authService.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, item := range *data {
		response.Data = append(response.Data, &pb.GetAuthDataResponse_Data{
			Id:       item.ID,
			Meta:     item.Meta,
			Login:    item.Login,
			Password: item.Password,
		})
	}
	return response, nil
}

func (s *GophKeeperServer) UpdateAuthData(ctx context.Context, in *pb.UpdateAuthDataRequest) (*pb.UpdateAuthDataResponse, error) {
	if in.Data == nil {
		return nil, nil
	}

	userID := ctx.Value(auth.UserIdContextKey).(string)
	data := models.AuthData{
		ID:       in.Data.Id,
		Meta:     in.Data.Meta,
		Login:    in.Data.Login,
		Password: in.Data.Password,
	}
	_, err := s.authService.Update(ctx, data, userID)
	if err != nil {
		return nil, err
	}
	s.sendNotifications(ctx, in.Data.Meta, "update", in.Data.Id)

	return &pb.UpdateAuthDataResponse{Result: true}, nil
}

func (s *GophKeeperServer) DeleteAuthData(ctx context.Context, in *pb.DeleteAuthDataRequest) (*pb.DeleteAuthDataResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := s.authService.Delete(ctx, in.Id, userID)
	if err != nil {
		return &pb.DeleteAuthDataResponse{Result: false}, err
	}
	s.sendNotifications(ctx, "authData", "delete", in.Id)
	return &pb.DeleteAuthDataResponse{Result: true}, err
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
