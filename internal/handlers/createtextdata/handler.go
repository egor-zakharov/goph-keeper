package createtextdata

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

func (h *Handler) Handle(ctx context.Context, in *pb.CreateTextDataRequest) (*pb.CreateTextDataResponse, error) {
	response := &pb.CreateTextDataResponse{}
	userID := ctx.Value(auth.UserIDContextKey).(string)
	if in.Data == nil {
		logger.Log().Sugar().Errorw("Create text data handler", "empty data error")
		return nil, status.Errorf(codes.InvalidArgument, "empty data error")
	}
	authData := models.TextData{
		Meta: in.Data.Meta,
		Text: in.Data.Text,
	}
	data, err := h.service.Create(ctx, authData, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Create text data handler", "create text service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	response.Id = data.ID
	h.notification.Send(ctx, notification.ProductText, notification.ActionCreate, response.Id)
	return response, nil
}
