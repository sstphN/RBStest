package services

import (
	"context"

	"go-wsl-project/internal/core/domain"
)

type ItemRepository interface {
	List(ctx context.Context) ([]domain.Item, error)
	Create(ctx context.Context, name string) (domain.Item, error)
}

type ItemService struct {
	repo ItemRepository
}

func NewItemService(r ItemRepository) *ItemService {
	return &ItemService{repo: r}
}

func (s *ItemService) List(ctx context.Context) ([]domain.Item, error) {
	return s.repo.List(ctx)
}

func (s *ItemService) Create(ctx context.Context, name string) (domain.Item, error) {
	return s.repo.Create(ctx, name)
}
