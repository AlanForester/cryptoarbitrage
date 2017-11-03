package helpers

import (
	"net/http"
	"time"
)

var HTTPClient = newHttpClient()

type HTTPClientParamsModel struct {
	url string
	method string
}

type HTTPClientModel struct {
	params HTTPClientParamsModel
	client http.Client
	Response http.Response
}

func (c *HTTPClientModel) SetHeaders(headers map[string][] string) HTTPClientModel {
	return *c
}

func (c*HTTPClientModel) SetTimeout(sec time.Duration)  {
	c.client.Timeout = sec
}

func (c *HTTPClientModel) Get(url string) HTTPClientModel {
	c.params.method = "GET"
	c.params.url = url
	c.query(c.params)
	return *c
}

func (c *HTTPClientModel) query(p HTTPClientParamsModel) HTTPClientModel {
	req, _ := http.NewRequest(c.params.method, "http://api.themoviedb.org/3/tv/popular", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent","Paw/3.1.4 (Macintosh; OS X/10.12.6) GCDHTTPRequest")
	resp, err := c.client.Do(req)

	if err != nil {
		return *c
	}

	defer resp.Body.Close()
	c.Response = *resp
	return *c
}

func newHttpClient() HTTPClientModel {
	c := new(HTTPClientModel)
	c.client = http.Client{Timeout: 10 * time.Second}
	return *c
}