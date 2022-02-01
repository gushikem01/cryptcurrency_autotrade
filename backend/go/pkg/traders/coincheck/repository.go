package coincheck

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/model"
)

// RepositoryPublic ...
type RepositoryPublic interface {
	Ticker(ctx context.Context) ([]*model.Ticker, error)
	TickerBtcJpy(ctx context.Context) (*model.Ticker, error)
}
