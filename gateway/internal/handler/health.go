package handler

import (
	"gateway/usecase"
	"net/http"
)

type HealthHandler struct {
	u *usecase.HealthUsecase
}

func NewHealthHandler(u *usecase.HealthUsecase) *HealthHandler {
	return &HealthHandler{
		u: u,
	}
}

func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := h.u.RespondHealth(ctx)

	respondJSON(w,r,resp)
}
