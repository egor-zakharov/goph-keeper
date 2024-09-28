package updateauthdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egor-zakharov/goph-keeper/internal/service/authdata"

	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
)

type Handler struct {
	service      authdata.Service
	notification notification.Service
}

func New(service authdata.Service, notification notification.Service) *Handler {
	return &Handler{
		service:      service,
		notification: notification,
	}
}

func (h *Handler) Handle(ctx context.Context, in *pb.UpdateAuthDataRequest) (*pb.UpdateAuthDataResponse, error) {
	if in.Data == nil {
		logger.Log().Sugar().Errorw("Update auth data handler", "empty data error")
		return &pb.UpdateAuthDataResponse{Result: false}, status.Errorf(codes.InvalidArgument, "empty data error")
	}

	userID := ctx.Value(auth.UserIDContextKey).(string)
	data := models.AuthData{
		ID:       in.Data.Id,
		Meta:     in.Data.Meta,
		Login:    in.Data.Login,
		Password: in.Data.Password,
	}
	_, err := h.service.Update(ctx, data, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Update auth data handler", "update auth data service", err)
		return &pb.UpdateAuthDataResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, notification.ProductAuth, notification.ActionUpdate, in.Data.Id)

	return &pb.UpdateAuthDataResponse{Result: true}, nil
}
