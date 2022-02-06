package public_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public"
	"go.uber.org/zap"
)

// TestTicker ...
func TestTicker(t *testing.T) {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	t.Run("Ticker get", func(t *testing.T) {
		tickerRes, err := r.Ticker(context.Background())
		if err != nil {
			t.Errorf("error:%s", err)
		}
		fmt.Println(tickerRes)
	})
}
