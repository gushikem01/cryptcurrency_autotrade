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
func (s *Service) Product(ctx context.Context) ([]*model.Products, error) {
	m, err := s.repo.Product(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}
