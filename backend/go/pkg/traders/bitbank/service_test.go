package bitbank_test

import (
	"context"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public"
	"go.uber.org/zap"
)

// initService ...
func initService() *bitbank.Service {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	s := bitbank.NewService(l, r)
	return s
}

// TestTicker ...
func TestTicker(t *testing.T) {
	s := initService()
	t.Run("Service test", func(t *testing.T) {
		s.Tickers(context.Background(), "btc_jpy")
	})
}
