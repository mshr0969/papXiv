package server

import (
	"gateway/handler"
	"net/http"
)

type Server struct {
	health *handler.HealthHandler
	paper  *handler.PaperHandler
}

func NewServer(h handler.Handlers) handler.ServerInterface {
	return &Server{
		health: h.Health,
		paper:  h.Paper,
	}
}
func (s *Server) HealthGet(w http.ResponseWriter, r *http.Request) {
	s.health.GetHealth(w, r)
}

func (s *Server) PaperDelete(w http.ResponseWriter, r *http.Request, paperId handler.PaperId) {
	s.paper.DeletePaper(w, r, paperId)
}

func (s *Server) PaperGet(w http.ResponseWriter, r *http.Request, paperId handler.PaperId) {
	s.paper.SelectPaper(w, r, paperId)
}

func (s *Server) PapersGet(w http.ResponseWriter, r *http.Request) {
	s.paper.ListPapers(w, r)
}

func (s *Server) PapersPost(w http.ResponseWriter, r *http.Request) {
	s.paper.CreatePaper(w, r)
}

func (s *Server) PaperPut(w http.ResponseWriter, r *http.Request, paperId handler.PaperId) {
	s.paper.UpdatePaper(w, r, paperId)
}

func (s *Server) SearchGet(w http.ResponseWriter, r *http.Request, params handler.SearchGetParams) {}
