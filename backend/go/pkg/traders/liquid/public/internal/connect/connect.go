package connect

import (
	"fmt"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/http_client"
)

const apiUrl string = "https://api.liquid.com/products/"

// Connection
type Connection struct {
}

// New ...
func New() *Connection {
	return &Connection{}
}

// Get ...
func (c *Connection) Get(path string, m map[string]string) ([]byte, error) {
	http := http_client.NewHTTP()
	status, ret, err := http.Get(apiUrl+path, m, nil)
	if err != nil || status >= 400 {
		return nil, fmt.Errorf("error:%v body:%v", err, ret)
	}
	return ret, nil
}
