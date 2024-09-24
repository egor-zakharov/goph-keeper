package uploadfile

import (
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

func (h *Handler) Handle(stream pb.GophKeeperServer_UploadFileServer) error {
	ctx := stream.Context()
	fileData, err := h.service.Add(ctx, stream)
	if err != nil {
		logger.Log().Sugar().Errorw("Upload file handler", "upload file service", err)
		return status.Errorf(codes.Internal, "internal error")
	}

	h.notification.Send(ctx, fileData.Meta, "upload", fileData.ID)
	return nil
}
