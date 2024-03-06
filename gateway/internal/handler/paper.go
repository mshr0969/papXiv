package handler

import "gateway/usecase"

type PaperHandler struct {
	u *usecase.PaperUsecase
}

func NewPaperHandler(u *usecase.PaperUsecase) *PaperHandler {
	return &PaperHandler{
		u: u,
	}
}
