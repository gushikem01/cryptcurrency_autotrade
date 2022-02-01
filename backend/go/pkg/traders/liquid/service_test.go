package liquid_test

import (
	"context"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public"
	"go.uber.org/zap"
)

// initService ...
func initService() *liquid.Service {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	s := liquid.NewService(l, r)
	return s
}

// TestService ...
func TestService(t *testing.T) {
	s := initService()
	t.Run("Service test", func(t *testing.T) {
		s.Product(context.Background())
	})
}
