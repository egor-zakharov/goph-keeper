package gettextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
		logger.Log().Sugar().Errorw("Get text data handler", "read text data service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
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
