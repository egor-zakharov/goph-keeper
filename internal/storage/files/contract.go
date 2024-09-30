package files

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=files

var (
	ErrConflict = errors.New("data conflict")
)

type Storage interface {
	Create(ctx context.Context, fileData models.FileData, userID string) (*models.FileData, error)
	Read(ctx context.Context, userID string) (*[]models.FileData, error)
	Delete(ctx context.Context, id string, userID string) error
}
