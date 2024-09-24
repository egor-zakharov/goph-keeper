package subcribetochanges

import (
	"context"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	"time"
)

type Handler struct {
	notification notification.Service
}

func New(notification notification.Service) *Handler {
	return &Handler{
		notification: notification,
	}
}

func (h *Handler) SubscribeToChanges(_ *pb.SubscribeToChangesRequest, stream pb.GophKeeper_SubscribeToChangesServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()
	h.notification.Add(ctx, stream)
	for {
		time.Sleep(time.Minute)
	}
}
