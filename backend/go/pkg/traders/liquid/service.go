package liquid

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public/model"
	"go.uber.org/zap"
)

// Service ...
type Service struct {
	logger *zap.Logger
	repo   RepositoryPublic
}

// NewService
func NewService(logger *zap.Logger, repo RepositoryPublic) *Service {
	return &Service{logger: logger, repo: repo}
}

// Product ...
func (s *Service) Products(ctx context.Context) ([]*model.Products, error) {
	m, err := s.repo.Products(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// ProductBtcJpy ...
func (s *Service) ProductBtcJpy(ctx context.Context) (*model.Products, error) {
	m, err := s.Products(ctx)
	if err != nil {
		return nil, err
	}

	b, err := s.btcJpy(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// btcJpy ...
func (s *Service) btcJpy(m []*model.Products) (*model.Products, error) {
	for _, v := range m {
		if v.CurrencyPairCode == "BTCJPY" {
			return v, nil
		}
	}
	return nil, nil
}
