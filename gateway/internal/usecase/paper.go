package usecase

import (
	"context"
	"fmt"
	"gateway/internal/domain"
	"gateway/internal/repository"
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

func (u *PaperUsecase) SelectPaper(ctx context.Context, paperId string) (*domain.Paper, error) {
	do, err := u.pr.SelectPaper(ctx, paperId)
	if err != nil {
		return nil, fmt.Errorf("error in u.pr.SelectPaper: %w", err)
	}
	return do, nil
}

func (u *PaperUsecase) UpdatePaper(ctx context.Context, do domain.Paper) error {
	if err := u.pr.UpdatePaper(ctx, do); err != nil {
		return fmt.Errorf("error in u.pr.UpdatePaper: %w", err)
	}
	return nil
}

func (u *PaperUsecase) DeletePaper(ctx context.Context, paperId string) error {
	if err := u.pr.DeletePaper(ctx, paperId); err != nil {
		return fmt.Errorf("error in u.pr.DeletePaper: %w", err)
	}
	return nil
}
