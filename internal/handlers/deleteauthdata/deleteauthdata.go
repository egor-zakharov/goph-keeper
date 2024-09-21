package deleteauthdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"

	"github.com/egor-zakharov/goph-keeper/internal/service/authdata"

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

func (h *Handler) Handle(ctx context.Context, in *pb.DeleteAuthDataRequest) (*pb.DeleteAuthDataResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := h.service.Delete(ctx, in.Id, userID)
	if err != nil {
		return &pb.DeleteAuthDataResponse{Result: false}, err
	}
	h.notification.Send(ctx, "authData", "delete", in.Id)
	return &pb.DeleteAuthDataResponse{Result: true}, err
}
