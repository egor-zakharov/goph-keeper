package deletetextdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (h *Handler) Handle(ctx context.Context, in *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	userID := ctx.Value(auth.UserIDContextKey).(string)
	err := h.service.Delete(ctx, in.Id, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Delete text data handler", "delete text data service", err)
		return &pb.DeleteTextDataResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, notification.ProductText, notification.ActionDelete, in.Id)
	return &pb.DeleteTextDataResponse{Result: true}, nil
}
