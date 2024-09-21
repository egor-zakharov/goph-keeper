package getcards

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
)

type Handler struct {
	cards cards.Service
}

func New(cards cards.Service) *Handler {
	return &Handler{
		cards: cards,
	}
}

func (h *Handler) Handle(ctx context.Context, _ *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	response := &pb.GetCardsResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	cards, err := h.cards.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, card := range *cards {
		response.Cards = append(response.Cards, &pb.GetCardsResponse_Card{
			Id:             card.ID,
			Number:         card.Number,
			ExpirationDate: card.ExpirationDate,
			HolderName:     card.HolderName,
			Cvv:            card.CVV,
		})
	}

	return response, nil
}
