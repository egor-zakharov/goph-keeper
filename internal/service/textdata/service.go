package textdata

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/textdata"
)

type service struct {
	storage textdata.Storage
}

func New(storage textdata.Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	return s.storage.Create(ctx, textData, userID)
}

func (s *service) Read(ctx context.Context, userID string) (*[]models.TextData, error) {
	return s.storage.Read(ctx, userID)
}

func (s *service) Update(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	return s.storage.Update(ctx, textData, userID)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.storage.Delete(ctx, id, userID)
}
