package model

import (
	"github.com/shopspring/decimal"
)

type Ticker struct {
	Last      decimal.Decimal `json: "last"`
	Bid       decimal.Decimal `json: "bid"`
	Ask       decimal.Decimal `json: "ask"`
	High      decimal.Decimal `json: "high"`
	Low       decimal.Decimal `json: "low"`
	Volume    decimal.Decimal `json: "volume"`
	Timestamp uint64          `json: "timestamp"`
}
