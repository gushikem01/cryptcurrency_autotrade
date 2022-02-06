package public

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public/internal/connect"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public/model"
	"go.uber.org/zap"
)

type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository
func NewRepository(l *zap.Logger) bitflayer.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}
}

// Ticker ...
func (repo *repo) Ticker(ctx context.Context) (*model.Ticker, error) {
	var wg sync.WaitGroup
	t := &model.Ticker{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ret, err := repo.conn.Get("ticker", nil)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		err = json.Unmarshal(ret, t)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
	}()
	wg.Wait()
	return t, nil
}
