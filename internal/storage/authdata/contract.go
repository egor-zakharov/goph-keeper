package authdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=authdata

var (
	ErrConflict = errors.New("data conflict")
)

type Storage interface {
	Create(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error)
	Read(ctx context.Context, userID string) (*[]models.AuthData, error)
	Update(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error)
	Delete(ctx context.Context, id string, userID string) error
}
