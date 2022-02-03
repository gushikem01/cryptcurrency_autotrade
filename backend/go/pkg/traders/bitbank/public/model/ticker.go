package model

type Ticker struct {
	Success int `json: "success"`
	Data    []struct {
		Pair      string `json: "pair"`
		Sell      string `json: "sell"`
		Buy       string `json: "buy"`
		Open      string `json: "open"`
		high      string `json: "hight"`
		Low       string `json: "low"`
		Last      string `json: "last"`
		Vol       string `json: "vol"`
		Timestamp int64  `json: "timestamp"`
	} `json: data`
}
