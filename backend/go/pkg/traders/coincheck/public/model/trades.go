package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Trades ...
type Trades struct {
	Data []struct {
		Id        int64           `json: "id"`
		Amount    decimal.Decimal `json: "amount"`
		Rate      decimal.Decimal `json: "rate"`
		Pair      string          `json: "pair"`
		OrderType string          `json: string`
		CreatedAt time.Time       `json: "created_at"`
	} `json: data`
}
