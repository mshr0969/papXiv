package usecase

import (
	"gateway/repository"
)

type PaperUsecase struct{
	pr repository.PaperRepository
}

func NewPaperUsecase(pr repository.PaperRepository) *PaperUsecase {
	return &PaperUsecase{
		pr: pr,
	}
}
