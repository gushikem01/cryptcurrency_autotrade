package btcbox

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox/public/model"
)

// RepositoryPublic ...
type RepositoryPublic interface {
	Tickers(ctx context.Context, pairs ...string) ([]*model.Ticker, error)
}
