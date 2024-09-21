package deletecard

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
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

func (s *Handler) Handle(ctx context.Context, in *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := s.cards.Delete(ctx, in.Id, userID)
	if err != nil {
		return &pb.DeleteCardResponse{Result: false}, err
	}
	s.notification.Send(ctx, "card", "delete", in.Id)
	return &pb.DeleteCardResponse{Result: true}, err
}
