package auth

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/auth"
)

type service struct {
	storage auth.Storage
}

func New(storage auth.Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error) {
	return s.storage.Create(ctx, authData, userID)
}

func (s *service) Read(ctx context.Context, userID string) (*[]models.AuthData, error) {
	return s.storage.Read(ctx, userID)
}

func (s *service) Update(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error) {
	return s.storage.Update(ctx, authData, userID)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.storage.Delete(ctx, id, userID)
}
