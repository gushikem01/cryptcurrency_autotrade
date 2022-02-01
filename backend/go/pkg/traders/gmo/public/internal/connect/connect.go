package connect

import (
	"fmt"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/http_client"
)

const apiURL = "https://api.coin.z.com/public"

// Connection
type Connection struct {
}

// New ...
func New() *Connection {
	return &Connection{}
}

// Get
func (c *Connection) Get(path string) ([]byte, error) {
	http := http_client.NewHTTP()
	status, ret, err := http.Get(apiURL+path, nil, nil)
	if err != nil || status >= 400 {
		return nil, fmt.Errorf("error:%v body:%v", err, ret)
	}
	return ret, nil

}
