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

func (h *PaperHandler) ListPapers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dos, err := h.u.ListPapers(ctx)
	if err != nil {
		handle(w, r, err)
		return
	}
	total := len(dos)

	resp := NewPaperListFromDomains(dos, total)
	respondJSON(w, r, resp)
}

func NewPaperItemFromDomain(do domain.Paper) PaperItem {
	return PaperItem{
		Id:         do.Id,
		Title: do.Title,
	}
}

func NewPaperListFromDomains(dos domain.Papers, total int) PaperList {
	rs := make([]PaperItem, 0, len(dos))
	for _, do := range dos {
		rs = append(rs, NewPaperItemFromDomain(do))
	}
	return PaperList{
		Total:   total,
		Papers:   &rs,
	}
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
