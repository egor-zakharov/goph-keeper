package updatecard

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
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

func (h *Handler) UpdateCard(ctx context.Context, in *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	if in.Card == nil {
		return nil, nil
	}
	userID := ctx.Value(auth.UserIdContextKey).(string)
	card := models.Card{
		ID:             in.Card.Id,
		Number:         in.Card.Number,
		ExpirationDate: in.Card.ExpirationDate,
		HolderName:     in.Card.HolderName,
		CVV:            in.Card.Cvv,
	}
	if !card.IsValidNumber() || !card.IsValidDate() || !card.IsValidNumber() {
		logger.Log().Sugar().Errorw("UpdateCard handler", "validation error")
		return &pb.UpdateCardResponse{Result: false}, status.Errorf(codes.InvalidArgument, "Incorrect card data")
	}
	_, err := h.cards.Update(ctx, card, userID)
	if err != nil {
		return nil, err
	}
	h.notification.Send(ctx, "card", "update", in.Card.Id)

	return &pb.UpdateCardResponse{Result: true}, nil
}
