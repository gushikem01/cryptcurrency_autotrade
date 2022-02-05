package model

// Ticker Struct
type Ticker struct {
	Pair      string `json: "pair"`
	Sell      string `json: "sell"`
	Buy       string `json: "buy"`
	Open      string `json: "open"`
	high      string `json: "hight"`
	Low       string `json: "low"`
	Last      string `json: "last"`
	Vol       string `json: "vol"`
	Timestamp int64  `json: "timestamp"`
}

// Res Struct
type Res struct {
	Success int      `json: "success"`
	Data    []Ticker `json: data`
}

// Change ...
func Change(v Ticker) *Ticker {
	m := &Ticker{}
	m.Buy = v.Buy
	m.Last = v.Last
	m.Low = v.Low
	m.Open = v.Open
	m.Pair = v.Pair
	m.Sell = v.Sell
	m.Timestamp = v.Timestamp
	m.Vol = v.Vol
	return m
}
