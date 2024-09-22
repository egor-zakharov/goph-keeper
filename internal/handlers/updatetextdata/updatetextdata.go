package updatetextdata

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

func (h *Handler) Handle(ctx context.Context, in *pb.UpdateConfTextDataRequest) (*pb.UpdateConfTextDataResponse, error) {
	if in.Data == nil {
		return nil, nil
	}

	userID := ctx.Value(auth.UserIdContextKey).(string)
	data := models.TextData{
		ID:   in.Data.Id,
		Meta: in.Data.Meta,
		Text: in.Data.Text,
	}
	_, err := h.service.Update(ctx, data, userID)
	if err != nil {
		return nil, err
	}
	h.notification.Send(ctx, in.Data.Meta, "update", in.Data.Id)

	return &pb.UpdateConfTextDataResponse{Result: true}, nil
}
