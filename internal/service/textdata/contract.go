package textdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=textdata

var ErrIncorrectData = errors.New("data incorrect")

type Service interface {
	Create(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error)
	Read(ctx context.Context, userID string) (*[]models.TextData, error)
	Update(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error)
	Delete(ctx context.Context, id string, userID string) error
}
