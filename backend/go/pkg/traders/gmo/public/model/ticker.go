package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Ticker ...
type Ticker struct {
	Data []struct {
		Ask       decimal.Decimal `json:"ask"`
		Bid       decimal.Decimal `json:"bid"`
		High      decimal.Decimal `json:"high"`
		Low       decimal.Decimal `json:"low"`
		Last      decimal.Decimal `json:"last"`
		Symbol    string          `json:"symbol"`
		Timestamp time.Time       `json:"timestamp"`
		Volume    decimal.Decimal `json:"volume"`
	}
}
