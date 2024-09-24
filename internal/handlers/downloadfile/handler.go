package downloadfile

import (
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/service/files"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service files.Service
}

func New(service files.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(in *pb.DownloadFileRequest, stream pb.GophKeeperServer_DownloadFileServer) error {
	err := h.service.Download(in, stream)
	if err != nil {
		logger.Log().Sugar().Errorw("Download file handler", "file service download", err)
		return status.Errorf(codes.Internal, "internal error")
	}
	return nil
}
