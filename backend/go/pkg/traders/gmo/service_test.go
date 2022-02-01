package gmo_test

import (
	"context"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public"
	"go.uber.org/zap"
)

// initService ...
func initService() *gmo.Service {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	s := gmo.NewService(l, r)
	return s
}

// TestService ...
func TestService(t *testing.T) {
	s := initService()
	t.Run("Service test", func(t *testing.T) {
		s.Ticker(context.Background())
	})
}
