package gettextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"

	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
)

type Handler struct {
	service textdata.Service
}

func New(service textdata.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx context.Context, _ *pb.GetConfTextDataRequest) (*pb.GetConfTextDataResponse, error) {
	response := &pb.GetConfTextDataResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	data, err := h.service.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, item := range *data {
		response.Data = append(response.Data, &pb.GetConfTextDataResponse_Data{
			Id:   item.ID,
			Meta: item.Meta,
			Text: item.Text,
		})
	}
	return response, nil
}
