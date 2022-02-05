package bitbank

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public/model"
	"go.uber.org/zap"
)

// Service ...
type Service struct {
	logger *zap.Logger
	repo   RepositoryPublic
}

// NewService ...
func NewService(logger *zap.Logger, repo RepositoryPublic) *Service {
	return &Service{logger: logger, repo: repo}
}

// Ticker ...
func (s *Service) Tickers(ctx context.Context, pairs ...string) ([]*model.Ticker, error) {
	m, err := s.repo.Tickers(ctx, pairs...)
	if err != nil {
		return nil, err
	}
	return m, nil
}
