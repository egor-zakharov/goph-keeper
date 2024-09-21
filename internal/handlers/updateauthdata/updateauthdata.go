package updateauthdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"

	"github.com/egor-zakharov/goph-keeper/internal/service/authdata"

	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
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

func (h *Handler) UpdateAuthData(ctx context.Context, in *pb.UpdateAuthDataRequest) (*pb.UpdateAuthDataResponse, error) {
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
	_, err := h.service.Update(ctx, data, userID)
	if err != nil {
		return nil, err
	}
	h.notification.Send(ctx, in.Data.Meta, "update", in.Data.Id)

	return &pb.UpdateAuthDataResponse{Result: true}, nil
}
