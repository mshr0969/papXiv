package usecase_test

import (
	"context"
	"gateway/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockPaperUsecase struct {
    mock.Mock
}

func (m *MockPaperUsecase) CreatePaper(ctx context.Context, do domain.Paper) error {
    args := m.Called(ctx, do)
    return args.Error(0)
}

func (m *MockPaperUsecase) ListPapers(ctx context.Context) (domain.Papers, error) {
    args := m.Called(ctx)
    return args.Get(0).(domain.Papers), args.Error(1)
}

func (m *MockPaperUsecase) SelectPaper(ctx context.Context, paperId string) (*domain.Paper, error) {
    args := m.Called(ctx, paperId)
    return args.Get(0).(*domain.Paper), args.Error(1)
}

func (m *MockPaperUsecase) UpdatePaper(ctx context.Context, do domain.Paper) error {
    args := m.Called(ctx, do)
    return args.Error(0)
}

func (m *MockPaperUsecase) DeletePaper(ctx context.Context, paperId string) error {
    args := m.Called(ctx, paperId)
    return args.Error(0)
}
