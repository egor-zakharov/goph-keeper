package cards

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=cards

var (
	ErrConflict = errors.New("data conflict")
)

type Storage interface {
	Create(ctx context.Context, card models.Card, userID string) (*models.Card, error)
	Read(ctx context.Context, userID string) (*[]models.Card, error)
	Update(ctx context.Context, card models.Card, userID string) (*models.Card, error)
	Delete(ctx context.Context, cardID string, userID string) error
}
