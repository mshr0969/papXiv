package handler

import (
	"encoding/json"
	"gateway/domain"
	"gateway/usecase"
	"net/http"
)

type PaperHandler struct {
	u *usecase.PaperUsecase
}

func NewPaperHandler(u *usecase.PaperUsecase) *PaperHandler {
	return &PaperHandler{
		u: u,
	}
}

func (h *PaperHandler) CreatePaper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PaperCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handle(w, r, err)
		return
	}

	do := req.toPaperDomain()

	err := h.u.CreatePaper(ctx, do)
	if err != nil {
		handle(w, r, err)
		return
	}

	respondCreated(w, r)
}

func (r PaperCreate) toPaperDomain() domain.Paper {
	return domain.Paper{
		Id:        r.Id,
		Published: r.Published,
		Subject:   r.Subject,
		Title:     r.Title,
		Url:       r.Url,
	}
}
