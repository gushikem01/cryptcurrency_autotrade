package bitflayer

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public/model"
)

// RepositoryPublic ...
type RepositoryPublic interface {
	Ticker(ctx context.Context) (*model.Ticker, error)
}
