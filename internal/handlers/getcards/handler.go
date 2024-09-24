package getcards

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		logger.Log().Sugar().Errorw("Get cards handler", "read card service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
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
