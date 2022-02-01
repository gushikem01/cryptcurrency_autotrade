package liquid

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public/model"
)

// RepositoryPublic ...
type RepositoryPublic interface {
	Product(ctx context.Context) ([]*model.Products, error)
}
