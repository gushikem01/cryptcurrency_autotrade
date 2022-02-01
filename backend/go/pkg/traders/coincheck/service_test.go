package coincheck_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public"
	"go.uber.org/zap"
)

// initService ...
func initService() *coincheck.Service {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	s := coincheck.NewService(l, r)
	return s
}

// TestService ...
func TestService(t *testing.T) {
	s := initService()
	t.Run("Service Test", func(t *testing.T) {
		ctx := context.Background()
		m, err := s.TickerBtcJpy(ctx)
		if err != nil {
			t.Fatalf(err.Error())
		}
		fmt.Println(m)
	})
}
