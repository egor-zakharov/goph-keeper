package signin

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/users"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	usersService users.Service
}

func New(usersService users.Service) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

func (s *Handler) Handle(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	response := &pb.SignInResponse{}

	user := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	if !user.IsValidLogin() || !user.IsValidPass() {
		logger.Log().Sugar().Errorw("Handle handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Login or password should not be empty")
	}

	usr, err := s.usersService.Login(ctx, user)

	if err != nil {
		logger.Log().Sugar().Errorw("Handle handler", "usersService login", err)
		return response, status.Errorf(codes.InvalidArgument, "Invalid login or password")
	}

	sessionID := uuid.New().String()
	JWTToken, err := auth.BuildJWTString(usr.UserID, sessionID)
	if err != nil {
		logger.Log().Sugar().Errorw("Handle handler", "build jwt", err)
		return response, status.Errorf(codes.Internal, "Can not build auth token")
	}

	response.Token = JWTToken
	return response, nil
}
