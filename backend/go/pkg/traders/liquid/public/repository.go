package public

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public/internal/connect"
	m "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public/model"
	"go.uber.org/zap"
)

var products = [...]int{5}

// repo ...
type repo struct {
	conn   *connect.Connection
	Logger *zap.Logger
}

// NewRepository ...
func NewRepository(l *zap.Logger) liquid.RepositoryPublic {
	conn := connect.New()
	return &repo{conn: conn, Logger: l}
}

// Products ...
func (repo *repo) Products(ctx context.Context) ([]*m.Products, error) {

	var wg sync.WaitGroup
	var p = []*m.Products{}

	for _, s := range products {
		wg.Add(1)

		go func(s int) {
			defer wg.Done()
			ret, err := repo.conn.Get(strconv.Itoa(s), nil)
			if err != nil {
				fmt.Println(err)
			}
			productRes := new(m.Products)
			err = json.Unmarshal(ret, productRes)
			if err != nil {
				fmt.Println(err)
			}

			p = append(p, productRes)
		}(s)
	}
	wg.Wait()
	return p, nil
}
