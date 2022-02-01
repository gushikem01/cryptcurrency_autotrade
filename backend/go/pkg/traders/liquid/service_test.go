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

// TestProducts ...
func TestProducts(t *testing.T) {
	s := initService()
	t.Run("Service test", func(t *testing.T) {
		s.Products(context.Background())
	})
}

func TestProductBtcJpy(t *testing.T) {
	s := initService()
	t.Run("Service test", func(t *testing.T) {
		s.ProductBtcJpy(context.Background())
	})
}
