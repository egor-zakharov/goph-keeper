package deletetextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"

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

func (h *Handler) Handle(ctx context.Context, in *pb.DeleteConfTextDataRequest) (*pb.DeleteConfTextDataResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := h.service.Delete(ctx, in.Id, userID)
	if err != nil {
		return &pb.DeleteConfTextDataResponse{Result: false}, err
	}
	h.notification.Send(ctx, "textData", "delete", in.Id)
	return &pb.DeleteConfTextDataResponse{Result: true}, err
}
