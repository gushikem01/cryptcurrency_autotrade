package coincheck

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/model"
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
func (s *Service) Ticker(ctx context.Context) ([]*model.Ticker, error) {
	m, err := s.repo.Ticker(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// TickerBtcJpy ...
func (s *Service) TickerBtcJpy(ctx context.Context) (*model.Ticker, error) {
	m, err := s.repo.TickerBtcJpy(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}
