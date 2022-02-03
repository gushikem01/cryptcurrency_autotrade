package public

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public/internal/connect"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public/model"
	"go.uber.org/zap"
)

type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository
func NewRepository(l *zap.Logger) bitbank.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}
}

// Ticker
func (repo *repo) Tickers(ctx context.Context) ([]*model.Ticker, error) {
	var wg sync.WaitGroup
	var t = []*model.Ticker{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ret, err := repo.conn.Get("tickers", nil)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		tickerRes := new(model.Ticker)
		err = json.Unmarshal(ret, tickerRes)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		t = append(t, tickerRes)
	}()
	wg.Wait()
	return t, nil
}
