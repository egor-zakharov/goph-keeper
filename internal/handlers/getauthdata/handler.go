package getauthdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egor-zakharov/goph-keeper/internal/service/authdata"

	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
)

type Handler struct {
	service authdata.Service
}

func New(service authdata.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx context.Context, _ *pb.GetAuthDataRequest) (*pb.GetAuthDataResponse, error) {
	response := &pb.GetAuthDataResponse{}
	userID := ctx.Value(auth.UserIdContextKey).(string)

	data, err := h.service.Read(ctx, userID)
	if err != nil {
		logger.Log().Sugar().Errorw("Get auth data handler", "read auth data service", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	for _, item := range *data {
		response.Data = append(response.Data, &pb.GetAuthDataResponse_Data{
			Id:       item.ID,
			Meta:     item.Meta,
			Login:    item.Login,
			Password: item.Password,
		})
	}
	return response, nil
}
