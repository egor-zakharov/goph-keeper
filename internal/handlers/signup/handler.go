package signup

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/users"
	usersStorage "github.com/egor-zakharov/goph-keeper/internal/storage/users"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
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

func (h *Handler) Handle(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	response := &pb.SignUpResponse{}

	user := models.User{
		Login:    in.Login,
		Password: in.Password,
	}
	if !user.IsValidLogin() || !user.IsValidPass() {
		logger.Log().Sugar().Errorw("Handle handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Login or password should not be empty")
	}
	createdUser, err := h.usersService.Register(ctx, user)

	if errors.Is(err, usersStorage.ErrConflict) {
		logger.Log().Sugar().Errorw("Handle handler", "usersService register", err)
		return response, status.Errorf(codes.InvalidArgument, "User with such login already exists")
	}

	if err != nil {
		logger.Log().Sugar().Errorw("Handle handler", "usersService register", err)
		return response, status.Errorf(codes.Internal, "Internal error")
	}

	sessionID := uuid.New().String()
	JWTToken, err := auth.BuildJWTString(createdUser.UserID, sessionID)

	if err != nil {
		logger.Log().Sugar().Errorw("Handle handler", "build jwt", err)
		return response, status.Errorf(codes.Internal, "Can not build auth token")
	}

	response.Token = JWTToken
	return response, nil

}
