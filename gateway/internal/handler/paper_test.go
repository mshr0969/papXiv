package handler

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain"
	usecase_test "gateway/internal/test/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/v3/assert"
)
func TestCreatePaper(t *testing.T) {
    u := new(usecase_test.MockPaperUsecase)

    h := NewPaperHandler(u)

    reqBody, _ := json.Marshal(PaperCreate{
        Id:        "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881",
        Published: "2024/01/01",
        Subject:   "physics",
        Title:     "test-title",
        Url:       "https://example.com",
    })
    req := httptest.NewRequest(http.MethodPost, "/papers", bytes.NewBuffer(reqBody))
    w := httptest.NewRecorder()

    u.On("CreatePaper", mock.Anything, mock.MatchedBy(func(p domain.Paper) bool {
        return p.Id == "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881" && p.Published == "2024/01/01" &&
            p.Subject == "physics" && p.Title == "test-title" && p.Url == "https://example.com"
    })).Return(nil)

    h.CreatePaper(w, req)

    resp := w.Result()
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    u.AssertExpectations(t)
}

func TestListPapers(t *testing.T) {
	u := new(usecase_test.MockPaperUsecase)

	h := NewPaperHandler(u)

	req := httptest.NewRequest(http.MethodGet, "/papers", nil)
	w := httptest.NewRecorder()

	u.On("ListPapers", mock.Anything).Return(domain.Papers{
		{
			Id:        "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881",
			Published: "2024/01/01",
			Subject:   "physics",
			Title:     "test-title",
			Url:       "https://example.com",
		},
	}, nil)

	h.ListPapers(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	u.AssertExpectations(t)
}

func TestSelectPaper(t *testing.T) {
	u := new(usecase_test.MockPaperUsecase)

	h := NewPaperHandler(u)

	req := httptest.NewRequest(http.MethodGet, "/papers/52132AF8-F5DF-461A-9CE1-8EEFC4DD4881", nil)
	w := httptest.NewRecorder()

	u.On("SelectPaper", mock.Anything, "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881").Return(&domain.Paper{
		Id:        "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881",
		Published: "2024/01/01",
		Subject:   "physics",
		Title:     "test-title",
		Url:       "https://example.com",
	}, nil)

	h.SelectPaper(w, req, "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881")

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	u.AssertExpectations(t)
}

func TestUpdatePaper(t *testing.T) {
	u := new(usecase_test.MockPaperUsecase)

	h := NewPaperHandler(u)

	reqBody, _ := json.Marshal(PaperUpdate{
		Published: strPtr("2024/01/01"),
		Subject:   strPtr("physics"),
		Title:     strPtr("test-title"),
	})
	req := httptest.NewRequest(http.MethodPut, "/papers/52132AF8-F5DF-461A-9CE1-8EEFC4DD4881", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	u.On("UpdatePaper", mock.Anything, mock.MatchedBy(func(p domain.Paper) bool {
		return p.Id == "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881" && p.Published == "2024/01/01" &&
			p.Subject == "physics" && p.Title == "test-title"
	})).Return(nil)

	h.UpdatePaper(w, req, "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881")

	resp := w.Result()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	u.AssertExpectations(t)
}

func TestDeletePaper(t *testing.T) {
	u := new(usecase_test.MockPaperUsecase)

	h := NewPaperHandler(u)

	req := httptest.NewRequest(http.MethodDelete, "/papers/52132AF8-F5DF-461A-9CE1-8EEFC4DD4881", nil)
	w := httptest.NewRecorder()

	u.On("DeletePaper", mock.Anything, "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881").Return(nil)

	h.DeletePaper(w, req, "52132AF8-F5DF-461A-9CE1-8EEFC4DD4881")

	resp := w.Result()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	u.AssertExpectations(t)
}

func strPtr(s string) *string {
	return &s
}
