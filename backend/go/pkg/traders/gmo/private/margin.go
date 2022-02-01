package private

import (
	"encoding/json"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/private/internal/connect"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/private/model"
)

// Service ...
type Service struct {
	apiKey    string
	secretKey string
}

// NewService ...
func NewService(apiKey string, secretKey string) *Service {
	return &Service{apiKey: apiKey, secretKey: secretKey}
}

// Margin ...
func (s *Service) Margin() (*model.Margin, error) {
	conn := connect.New(s.apiKey, s.secretKey)
	ret, err := conn.Get("/v1/account/margin")
	if err != nil {
		return nil, err
	}

	marginRes := new(model.Margin)
	err = json.Unmarshal(ret, marginRes)
	if err != nil {
		return nil, err
	}
	return marginRes, nil

}
