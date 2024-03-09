package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/internal/domain"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request, err error) {
	pd := new(ProblemDetail)

	switch {
	case errors.Is(err, domain.ErrNonExistentPaper):
		// 404: not found
		pd.ErrorCode = "not_found"
		pd.Status = http.StatusNotFound
		pd.Title = "Not Found"
	default:
		// 500: internal server error
		pd.ErrorCode = "internal_server_error"
		pd.Status = http.StatusInternalServerError
		pd.Title = "Internal Server Error"
	}

	pd.Instance = r.RequestURI
	pd.Detail = err.Error()

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")

	b, err := json.Marshal(pd)
	if err != nil {
		log.Printf("encode response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(pd.Status)

	if _, err := fmt.Fprintf(w, "%s", b); err != nil {
		log.Printf("write response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func respondJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("encode response error: %v", err)
		handle(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprintf(w, "%s", b); err != nil {
		log.Printf("write response error: %v", err)
		handle(w, r, err)
		return
	}
}

func respondCreated(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Created successfully"})
}

func respondNoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
