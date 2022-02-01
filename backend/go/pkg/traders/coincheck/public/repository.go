package public

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/internal/connect"
	m "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/model"
	"go.uber.org/zap"
)

var pairs = [...]string{"btc_jpy", "plt_jpy"}

type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository ...
func NewRepository(l *zap.Logger) coincheck.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}
}

// Ticker ...
func (repo *repo) Ticker(ctx context.Context) ([]*m.Ticker, error) {
	var wg sync.WaitGroup
	var t = []*m.Ticker{}

	for _, s := range pairs {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()

			tickerRes, err := repo.get(s)
			if err != nil {
				repo.Logger.Error(err.Error())
			}
			t = append(t, tickerRes)
		}(s)
	}
	wg.Wait()
	return t, nil
}

// TickerBtcJpy ...
func (repo *repo) TickerBtcJpy(ctx context.Context) (*m.Ticker, error) {
	var wg sync.WaitGroup
	t := &m.Ticker{}

	s := pairs[0]

	wg.Add(1)
	go func(s string) {
		defer wg.Done()
		tickerRes, err := repo.get(s)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		t = tickerRes
	}(s)
	wg.Wait()

	return t, nil
}

// get ...
func (repo *repo) get(s string) (*m.Ticker, error) {
	ret, err := repo.conn.Get(fmt.Sprintf("/ticker?pair=%v", s), nil)
	if err != nil {
		return nil, err
	}
	tickerRes := new(m.Ticker)
	err = json.Unmarshal(ret, tickerRes)
	if err != nil {
		return nil, err
	}
	return tickerRes, nil
}
