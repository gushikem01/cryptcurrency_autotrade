package connect

import (
	"fmt"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/http_client"
)

const apiURL string = "https://www.btcbox.co.jp/api/v1/"

// Connection ...
type Connection struct {
}

// Connection ...
func New() *Connection {
	return &Connection{}
}

// Get ...
func (c *Connection) Get(path string, m map[string]string) ([]byte, error) {
	http := http_client.NewHTTP()
	status, ret, err := http.Get(apiURL+path, m, nil)
	if err != nil || status >= 400 {
		return nil, fmt.Errorf("error: %v body:%v", err, ret)
	}
	return ret, nil
}
