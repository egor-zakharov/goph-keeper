package cards

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/cards"
)

type service struct {
	storage cards.Storage
}

func New(storage cards.Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, card models.Card, userID string) (*models.Card, error) {
	return s.storage.Create(ctx, card, userID)
}

func (s *service) Read(ctx context.Context, userID string) (*[]models.Card, error) {
	return s.storage.Read(ctx, userID)
}

func (s *service) Update(ctx context.Context, card models.Card, userID string) (*models.Card, error) {
	return s.storage.Update(ctx, card, userID)
}

func (s *service) Delete(ctx context.Context, cardID string, userID string) error {
	return s.storage.Delete(ctx, cardID, userID)
}
