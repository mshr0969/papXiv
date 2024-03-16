package handler

import (
	"encoding/json"
	"gateway/internal/domain"
	"gateway/internal/usecase"
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

func (h *PaperHandler) SelectPaper(w http.ResponseWriter, r *http.Request, paperId string) {
	ctx := r.Context()

	do, err := h.u.SelectPaper(ctx, paperId)
	if err != nil {
		handle(w, r, err)
		return
	}

	resp := NewPaperGetFromDomain(*do)
	respondJSON(w, r, resp)
}

func (h *PaperHandler) UpdatePaper(w http.ResponseWriter, r *http.Request, paperId string) {
	ctx := r.Context()

	var req PaperUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handle(w, r, err)
		return
	}

	do := req.toPaperDomain(paperId)

	if err := h.u.UpdatePaper(ctx, do); err != nil {
		handle(w, r, err)
		return
	}
}

func (h *PaperHandler) DeletePaper(w http.ResponseWriter, r *http.Request, paperId string) {
	ctx := r.Context()

	if err := h.u.DeletePaper(ctx, paperId); err != nil {
		handle(w, r, err)
		return
	}

	respondNoContent(w, r)
}

func NewPaperItemFromDomain(do domain.Paper) PaperItem {
	return PaperItem{
		Id:    do.Id,
		Title: do.Title,
	}
}

func NewPaperListFromDomains(dos domain.Papers, total int) PaperList {
	rs := make([]PaperItem, 0, len(dos))
	for _, do := range dos {
		rs = append(rs, NewPaperItemFromDomain(do))
	}
	return PaperList{
		Total:  total,
		Papers: &rs,
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

func (r PaperUpdate) toPaperDomain(paperId string) domain.Paper {
	var published string
	if r.Published != nil {
		published = *r.Published
	}
	var subject string
	if r.Subject != nil {
		subject = *r.Subject
	}
	var title string
	if r.Title != nil {
		title = *r.Title
	}
	var url string
	if r.Url != nil {
		url = *r.Url
	}
	return domain.Paper{
		Id:        paperId,
		Published: published,
		Subject:   subject,
		Title:     title,
		Url:       url,
	}
}

func NewPaperGetFromDomain(do domain.Paper) PaperGet {
	return PaperGet{
		Id:        do.Id,
		Published: do.Published,
		Subject:   do.Subject,
		Title:     do.Title,
		Url:       do.Url,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
}
