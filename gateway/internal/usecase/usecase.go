package usecase

import (
	"context"
	"gateway/internal/domain"
)

type Usecases struct {
	Health *HealthUsecase
	Paper  *PaperUsecase
}

type PaperUsecaseInterface interface {
    CreatePaper(ctx context.Context, do domain.Paper) error
    ListPapers(ctx context.Context) (domain.Papers, error)
    SelectPaper(ctx context.Context, paperId string) (*domain.Paper, error)
    UpdatePaper(ctx context.Context, do domain.Paper) error
    DeletePaper(ctx context.Context, paperId string) error
}
