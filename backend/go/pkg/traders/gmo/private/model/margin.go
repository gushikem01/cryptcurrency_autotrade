package model

import "github.com/shopspring/decimal"

// Margin 余力情報
type Margin struct {
	Status int `json: "status"`
	Data   struct {
		ActualProfitLoss decimal.Decimal `json:"actualProfitLoss"`
		AvailableAmout   decimal.Decimal `json:"availableAmount"`
		Margin           decimal.Decimal `json:"margin"`
		MarginCallStatus decimal.Decimal `json:"marginCallStatus"`
		MarginRatio      decimal.Decimal `json:"marginRatio"`
		ProfitLoss       decimal.Decimal `json:"profitLoss"`
	} `json: "data"`
}
