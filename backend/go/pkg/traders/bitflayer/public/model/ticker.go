package model

import "github.com/shopspring/decimal"

// Ticker ...
type Ticker struct {
	ProductCode     string          `json:"product_code"`
	State           string          `json:"state"`
	Timestamp       string          `json:"timestamp"`
	TickID          int             `json:"tick_id"`
	BestBid         decimal.Decimal `json:"best_bid"`
	BestAsk         decimal.Decimal `json:"best_ask"`
	BestBidSize     decimal.Decimal `json:"best_bid_size"`
	BestAskSize     decimal.Decimal `json:"best_ask_size"`
	TotalBidDepth   decimal.Decimal `json:"total_bid_depth"`
	TotalAskDepth   decimal.Decimal `json:"total_ask_depth"`
	MarketBidSize   decimal.Decimal `json:"market_bid_size"`
	MarketAskSize   decimal.Decimal `json:"market_ask_size"`
	Ltp             decimal.Decimal `json:"ltp"`
	Volume          decimal.Decimal `json:"volume"`
	VolumeByProduct decimal.Decimal `json:"volume_by_product"`
}
