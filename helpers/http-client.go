package helpers

import (
	"net/http"
	"io/ioutil"
	"time"
)

var HTTPClient = new(HTTPClientModel)

type HTTPClientParamsModel struct {
	url string
	method string
	headers map[string][] string
}

type HTTPClientModel struct {
	params HTTPClientParamsModel
	client *http.Client
	Response *http.Response
}

func (c *HTTPClientModel) SetHeaders(headers map[string][] string) HTTPClientModel {
	c.params.headers = headers
	return *c
}

func (c*HTTPClientModel) SetTimeout(sec time.Duration)  {
	c.client.Timeout = sec
}

func (c *HTTPClientModel) Get(url string) http.Response {
	c.params.method = "GET"
	c.params.url = url
	c.query(c.params)
	return *c.Response
}

func (c *HTTPClientModel) query(p HTTPClientParamsModel) HTTPClientModel {
	client := &http.Client{}

	req, _ := http.NewRequest(c.params.method, "http://api.themoviedb.org/3/tv/popular", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return *c
	}

	defer resp.Body.Close()
	c.Response = resp
	return *c
}

func init()  {
	HTTPClient = new(HTTPClientModel)
	HTTPClient.client = &http.Client{Timeout: 10 * time.Second}
}