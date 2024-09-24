package createcard

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

func (h *Handler) Handle(ctx context.Context, in *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	response := &pb.CreateCardResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)
	if in.Card == nil {
		logger.Log().Sugar().Errorw("Create card handler", "empty data error")
		return nil, status.Errorf(codes.InvalidArgument, "empty data error")
	}

	card := models.Card{
		Number:         in.Card.Number,
		ExpirationDate: in.Card.ExpirationDate,
		HolderName:     in.Card.HolderName,
		CVV:            in.Card.Cvv,
	}

	if !card.IsValidNumber() || !card.IsValidDate() || !card.IsValidNumber() {
		logger.Log().Sugar().Errorw("Handle handler", "validation error")
		return response, status.Errorf(codes.InvalidArgument, "Incorrect card data")
	}

	createdCard, err := h.cards.Create(ctx, card, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Create card handler", "create card service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response.CardID = createdCard.ID
	h.notification.Send(ctx, "card", "create", createdCard.ID)
	return response, nil
}
