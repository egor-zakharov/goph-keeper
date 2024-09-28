package deletecard

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	cards        cards.Service
	notification notification.Service
}

func New(cards cards.Service, notification notification.Service) *Handler {
	return &Handler{
		cards:        cards,
		notification: notification,
	}
}

func (h *Handler) Handle(ctx context.Context, in *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	userID := ctx.Value(auth.UserIDContextKey).(string)
	err := h.cards.Delete(ctx, in.Id, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Delete card handler", "delete card service", err)
		return &pb.DeleteCardResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, notification.ProductCard, notification.ActionDelete, in.Id)
	return &pb.DeleteCardResponse{Result: true}, nil
}
