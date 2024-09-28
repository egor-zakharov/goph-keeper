package updatetextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
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

func (h *Handler) Handle(ctx context.Context, in *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	if in.Data == nil {
		logger.Log().Sugar().Errorw("Update text data handler", "empty data error")
		return &pb.UpdateTextDataResponse{Result: false}, status.Errorf(codes.InvalidArgument, "empty data error")
	}

	userID := ctx.Value(auth.UserIDContextKey).(string)
	data := models.TextData{
		ID:   in.Data.Id,
		Meta: in.Data.Meta,
		Text: in.Data.Text,
	}
	_, err := h.service.Update(ctx, data, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Update text data handler", "update text data service", err)
		return &pb.UpdateTextDataResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, notification.ProductText, notification.ActionUpdate, in.Data.Id)

	return &pb.UpdateTextDataResponse{Result: true}, nil
}
