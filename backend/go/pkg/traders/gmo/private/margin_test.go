package private_test

import (
	"fmt"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/private"
)

func initService() *private.Service {
	// TODO: env
	s := private.NewService("", "")
	return s
}

func TestMargin(t *testing.T) {
	s := initService()
	t.Run("Margin get", func(t *testing.T) {
		marginRes, err := s.Margin()
		if err != nil {
			t.Errorf("error:%s", err)
		}
		fmt.Println(marginRes)
	})
}
