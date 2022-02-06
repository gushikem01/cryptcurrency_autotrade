package public_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox/public"
	"go.uber.org/zap"
)

func Test(t *testing.T) {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	t.Run("Ticker get", func(t *testing.T) {
		tickerRes, err := r.Tickers(context.Background(), "btc_jpy")
		if err != nil {
			t.Errorf("error:%s", err)
		}
		fmt.Println(tickerRes)
	})
}
