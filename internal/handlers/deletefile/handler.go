package deletefile

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/files"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service      files.Service
	notification notification.Service
}

func New(service files.Service, notification notification.Service) *Handler {
	return &Handler{
		service:      service,
		notification: notification,
	}
}

func (h *Handler) Handle(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	err := h.service.Delete(ctx, in.Id, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Delete file handler", "delete file service", err)
		return &pb.DeleteFileResponse{Result: false}, status.Errorf(codes.Internal, "internal error")
	}
	h.notification.Send(ctx, "file", "delete", in.Id)
	return &pb.DeleteFileResponse{Result: true}, nil
}
