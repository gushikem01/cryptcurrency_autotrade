package public

import (
	"context"
	"encoding/json"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public/internal/connect"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public/model"
	"go.uber.org/zap"
)

// repo ...
type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository ...
func NewRepository(l *zap.Logger) gmo.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}

}

// Ticker ...
func (repo *repo) Ticker(ctx context.Context) (*model.Ticker, error) {
	ret, err := repo.conn.Get("/v1/ticker")
	if err != nil {
		return nil, err
	}
	tickerRes := new(model.Ticker)
	err = json.Unmarshal(ret, tickerRes)
	if err != nil {
		return nil, err
	}
	return tickerRes, nil
}
