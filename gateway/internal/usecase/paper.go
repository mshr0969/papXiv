package usecase

import (
	"context"
	"fmt"
	"gateway/domain"
	"gateway/repository"
)

type PaperUsecase struct {
	pr repository.PaperRepositoryInterface
}

func NewPaperUsecase(pr repository.PaperRepositoryInterface) *PaperUsecase {
	return &PaperUsecase{
		pr: pr,
	}
}

func (u *PaperUsecase) CreatePaper(ctx context.Context, do domain.Paper) error {
	if err := u.pr.CreatePaper(ctx, do); err != nil {
		return fmt.Errorf("error in u.pr.CreatePaper: %w", err)
	}
	return nil
}

func (u *PaperUsecase) ListPapers(ctx context.Context) (domain.Papers, error) {
	dos, err := u.pr.ListPapers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in u.pr.ListPapers; %w", err)
	}
	return dos, nil
}
