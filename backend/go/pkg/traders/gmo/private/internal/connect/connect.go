package connect

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/http_client"
)

// Connection ...
type Connection struct {
	apiKey    string
	secretKey string
}

// New ...
func New(apiKey, secretKey string) *Connection {
	return &Connection{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Get ...
func (c *Connection) Get(path string) ([]byte, error) {
	http := http_client.NewHTTP()

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	text := timestamp + "GET" + path
	hc := hmac.New(sha256.New, []byte(c.secretKey))
	hc.Write([]byte(text))
	sign := hex.EncodeToString(hc.Sum(nil))

	m := map[string]string{
		"API-KEY":       c.apiKey,
		"API-TIMESTAMP": timestamp,
		"API-SIGN":      sign,
	}
	status, ret, err := http.Get("https://api.coin.z.com/private"+path, m, nil)
	if err != nil || status >= 400 {
		return nil, fmt.Errorf("error:%v body:%v", err, ret)
	}
	return ret, nil
}
