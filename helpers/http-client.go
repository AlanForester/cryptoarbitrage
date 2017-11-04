package helpers

import (
	"net/http"
	"time"
	"log"
	"io/ioutil"
)
// Инициализация при каждом вызове
var HTTPClient = newHTTPClient()
// var HTTPClient = newHTTPClient()
// var HTTPClient HTTPClientModel

type HTTPClientModel struct {
	params HTTPClientParamsModel
	client http.Client
	response HTTPClientResponse
}

type HTTPClientParamsModel struct {
	url string
	method string
}

type HTTPClientResponse struct {
	Body []byte
	StatusCode int
	Headers map[string][]string
	Error error
}

func (c *HTTPClientModel) SetHeaders(headers map[string][]string) *HTTPClientModel {
	return c
}

func (c *HTTPClientModel) SetTimeout(sec time.Duration) *HTTPClientModel  {
	c.client.Timeout = sec
	return c
}

func (c *HTTPClientModel) Get(url string) *HTTPClientResponse {
	c.params.method = "GET"
	c.params.url = url
	c.query(c.params)
	return c.query(c.params)
}

func (c *HTTPClientModel) query(p HTTPClientParamsModel) *HTTPClientResponse {
	req, _ := http.NewRequest(c.params.method, c.params.url, nil)
	req.Header.Add("User-Agent","Paw/3.1.4 (Macintosh; OS X/10.12.6) GCDHTTPRequest")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	c.response.StatusCode = resp.StatusCode
	c.response.Headers = resp.Header
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		c.response.Error = err
		return nil
	}
	c.response.Body = body
	return &c.response
}

// TODO: Провести тесты с обоими вариантами
// Инициализация при каждом вызове
func newHTTPClient() *HTTPClientModel {
	print("324234\n")
	c := new(HTTPClientModel)
	c.client = http.Client{Timeout: 10 * time.Second}
	return c
}

// Инициализация при загрузке библиотеки
//func init() {
//	print("324234\n")
//	c := new(HTTPClientModel)
//	c.client = http.Client{Timeout: 10 * time.Second}
//	HTTPClient = *c
//}