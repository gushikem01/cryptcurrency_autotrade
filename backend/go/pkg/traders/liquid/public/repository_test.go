package public_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public"
	"go.uber.org/zap"
)

// TestProducts ...
func TestProducts(t *testing.T) {
	l, _ := zap.NewDevelopment()
	r := public.NewRepository(l)
	t.Run("Products get", func(t *testing.T) {
		productsRes, err := r.Products(context.Background())
		if err != nil {
			t.Errorf("error: %s", err)
		}
		fmt.Println(productsRes)
	})
}
