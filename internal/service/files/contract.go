package files

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=files

type Service interface {
	Add(ctx context.Context, stream pb.GophKeeper_UploadFileServer) (*models.FileData, error)
	Read(ctx context.Context, userID string) (*[]models.FileData, error)
	Delete(ctx context.Context, id string, userID string) error
	Download(in *pb.DownloadFileRequest, stream pb.GophKeeper_DownloadFileServer) error
}
