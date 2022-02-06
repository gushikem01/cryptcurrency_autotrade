package model

// Ticker ...
type Ticker struct {
	High int     `json:"high"`
	Low  int     `json:"low"`
	Buy  int     `json:"buy"`
	Sell int     `json:"sell"`
	Last int     `json:"last"`
	Vol  float64 `json:"vol"`
}

// Res ...
type Res struct {
	BtcJpy Ticker `json:"BTC_JPY"`
	BchJpy Ticker `json:"BCH_JPY"`
	LtcJpy Ticker `json:"LTC_JPY"`
	EthJpy Ticker `json:"ETH_JPY"`
}

// Change ...
func Change(v Ticker) *Ticker {
	m := &Ticker{}
	m.High = v.High
	m.Buy = v.Buy
	m.Last = v.Last
	m.Low = v.Low
	m.Vol = v.Vol
	m.Sell = v.Sell
	return m
}
