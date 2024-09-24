package getfiles

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
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

func (h *Handler) Handle(ctx context.Context, _ *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	response := &pb.GetFilesResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	files, err := h.service.Read(ctx, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Get files handler", "read files service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	for _, file := range *files {
		response.Files = append(response.Files, &pb.GetFilesResponse_File{
			Id:   file.ID,
			Name: file.Name,
			Meta: file.Meta,
		})
	}

	return response, nil
}
