package http_client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// AUTH 認証ヘッダKey
const AUTH = "Authorization"

// URLInfo URL情報
type URLInfo struct {
	URL    string            `json: "url"`
	Header map[string]string `json: "header"`
}

// HTTP HTTPクライアント
type HTTP struct {
	Client http.Client
}

// NewHTTP HTTPクライアントコンストラクタ
func NewHTTP() *HTTP {
	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	h := &HTTP{Client: *client}
	return h
}

// Get Getリクエスト
func (h *HTTP) Get(url string, header map[string]string, body io.Reader) (int, []byte, error) {
	res, err := h.request(http.MethodGet, url, header, body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()
	// バイト変換
	b, err := ioutil.ReadAll(res.Body)
	return res.StatusCode, b, err
}

// Post Postリクエスト
func (h *HTTP) Post(url string, header map[string]string, body io.Reader) (int, []byte, error) {
	res, err := h.request(http.MethodPost, url, header, body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()

	// バイト変換
	b, err := ioutil.ReadAll(res.Body)
	return res.StatusCode, b, err

}

// Put Putリクエスト
func (h *HTTP) Put(url string, header map[string]string, body io.Reader) (int, []byte, error) {
	res, err := h.request(http.MethodPut, url, header, body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()

	// バイト変換
	b, err := ioutil.ReadAll(res.Body)
	return res.StatusCode, b, err
}

// Delete Deleteリクエスト
func (h *HTTP) Delete(url string, header map[string]string, body io.Reader) (int, []byte, error) {
	res, err := h.request(http.MethodDelete, url, header, body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	defer res.Body.Close()

	// バイト変換
	b, err := ioutil.ReadAll(res.Body)
	return res.StatusCode, b, err
}

// request httpリクエスト
func (h *HTTP) request(method, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// デフォルトヘッダー
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// ヘッダー設定
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// リクエスト
	res, err := h.Client.Do(req)
	if res.StatusCode >= 400 {
		fmt.Errorf("HTTP通信エラー")
		fmt.Println(err)
		// h.logger.Error("HTTP通信エラー",
		// 	zap.Error(err),
		// 	zap.Any("HTTPPayload", &zapdriver.HTTPPayload{
		// 		RequestMethod: method,
		// 		RequestURL:    url,
		// 		Status:        res.StatusCode,
		// 		Protocol:      res.Proto,
		// 		RequestSize:   strconv.FormatInt(res.Request.ContentLength, 10),
		// 		ResponseSize:  strconv.FormatInt(res.ContentLength, 10),
		// 	}),
		// )
	} else {
		fmt.Println("HTTP通信完了")
		// h.logger.Info("HTTP通信完了", zap.Object("HTTPPayload", &zapdriver.HTTPPayload{
		// 	RequestMethod: method,
		// 	RequestURL:    url,
		// 	Status:        res.StatusCode,
		// 	Protocol:      res.Proto,
		// 	RequestSize:   strconv.FormatInt(res.Request.ContentLength, 10),
		// 	ResponseSize:  strconv.FormatInt(res.ContentLength, 10),
		// }))
	}
	return res, err
}
