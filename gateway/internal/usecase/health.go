package usecase

import (
	"context"
)

type HealthUsecase struct{}

func NewHealthUsecase() *HealthUsecase {
	return &HealthUsecase{}
}

func (h *HealthUsecase) RespondHealth(ctx context.Context) map[string]string {
	return map[string]string{
		"message": "ok",
	}
}
