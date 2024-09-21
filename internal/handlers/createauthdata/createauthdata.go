package createauthdata

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

func (h *Handler) Handle(ctx context.Context, in *pb.CreateAuthDataRequest) (*pb.CreateAuthDataResponse, error) {
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
	data, err := h.service.Create(ctx, authData, userID)
	if err != nil {
		return nil, err
	}
	response.Id = data.ID
	h.notification.Send(ctx, authData.Meta, "create", response.Id)
	return response, nil
}
