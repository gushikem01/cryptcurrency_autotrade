package public

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox/public/internal/connect"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox/public/model"
	"go.uber.org/zap"
)

type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository ...
func NewRepository(l *zap.Logger) btcbox.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}
}

// Tickers ...
func (repo *repo) Tickers(ctx context.Context, pairs ...string) ([]*model.Ticker, error) {
	var wg sync.WaitGroup
	var t = []*model.Ticker{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// API実行
		ret, err := repo.conn.Get("tickers", nil)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		tickerRes := new(model.Res)
		err = json.Unmarshal(ret, tickerRes)
		if err != nil {
			repo.Logger.Error(err.Error())
		}
		// 絞り込み
		t = repo.filter(ctx, tickerRes, pairs...)
	}()
	wg.Wait()
	return t, nil
}

// filter ...
func (repo *repo) filter(ctx context.Context, res *model.Res, pairs ...string) []*model.Ticker {
	var t = []*model.Ticker{}
	for _, v := range pairs {
		if v == "btc_jpy" {
			m := model.Change(res.BtcJpy)
			t = append(t, m)
		}
		if v == "bch_jpy" {
			m := model.Change(res.BchJpy)
			t = append(t, m)
		}
		if v == "eth_jpy" {
			m := model.Change(res.EthJpy)
			t = append(t, m)
		}
		if v == "ltc_jpy" {
			m := model.Change(res.LtcJpy)
			t = append(t, m)
		}
	}
	return t
}
