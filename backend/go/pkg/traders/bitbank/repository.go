package bitbank

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public/model"
)

// RepositoryPublic ...
type RepositoryPublic interface {
	Tickers(ctx context.Context) ([]*model.Ticker, error)
}
