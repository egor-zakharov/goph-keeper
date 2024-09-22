package createtextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"

	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
)

type Handler struct {
	service      textdata.Service
	notification notification.Service
}

func New(service textdata.Service, notification notification.Service) *Handler {
	return &Handler{
		service:      service,
		notification: notification,
	}
}

func (h *Handler) Handle(ctx context.Context, in *pb.CreateConfTextDataRequest) (*pb.CreateConfTextDataResponse, error) {
	response := &pb.CreateConfTextDataResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)
	if in.Data == nil {
		return nil, nil
	}
	authData := models.TextData{
		Meta: in.Data.Meta,
		Text: in.Data.Text,
	}
	data, err := h.service.Create(ctx, authData, userID)
	if err != nil {
		return nil, err
	}
	response.Id = data.ID
	h.notification.Send(ctx, authData.Meta, "create", response.Id)
	return response, nil
}
