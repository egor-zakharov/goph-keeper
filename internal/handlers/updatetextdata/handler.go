package updatetextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
)

type Handler struct {
	service      textdata.Service
	notification notification.Service
}

func New(service textdata.Service, notification notification.Service) *Handler {
	return &Handler{
		service:      service,
		notification: notification,
	}
}

func (h *Handler) Handle(ctx context.Context, in *pb.UpdateConfTextDataRequest) (*pb.UpdateConfTextDataResponse, error) {
	if in.Data == nil {
		logger.Log().Sugar().Errorw("Update text data handler", "empty data error")
		return &pb.UpdateConfTextDataResponse{Result: false}, status.Errorf(codes.InvalidArgument, "empty data error")
	}

	userID := ctx.Value(auth.UserIdContextKey).(string)
	data := models.TextData{
		ID:   in.Data.Id,
		Meta: in.Data.Meta,
		Text: in.Data.Text,
	}
	_, err := h.service.Update(ctx, data, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Update text data handler", "update text data service", err)
		return &pb.UpdateConfTextDataResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, in.Data.Meta, "update", in.Data.Id)

	return &pb.UpdateConfTextDataResponse{Result: true}, nil
}
