package repository

import (
	"context"
	"gateway/domain"
)

type Repositories struct {
	Paper PaperRepositoryInterface
}

type PaperRepositoryInterface interface {
	ListPapers(ctx context.Context) (domain.Papers, error)
	SelectPaper(ctx context.Context, paperID string) (*domain.Paper, error)
	CreatePaper(ctx context.Context, do domain.Paper) error
	UpdatePaper(ctx context.Context, paperID string) error
	DeletePaper(ctx context.Context, paperID string) error
}
