package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/padmoney/enotas-go/config"
)

type Client interface {
	AddHeader(key, value string)
	Delete(url string) Response
	Get(url string) Response
	Post(url string, body []byte) Response
	Put(url string, body []byte) Response
	Send(method, url string, body io.Reader) Response
}

type client struct {
	credentials config.Credentials
	headers     map[string]string
}

func NewClient(credentials config.Credentials) *client {
	c := &client{
		credentials: credentials,
		headers:     make(map[string]string),
	}
	c.setAuthorization()
	return c
}

func (r *client) Delete(url string) Response {
	return r.Send(http.MethodDelete, url, nil)
}

func (r *client) Get(url string) Response {
	return r.Send(http.MethodGet, url, nil)
}

func (r *client) Post(url string, body []byte) Response {
	r.AddHeader("Content-Type", "application/json")
	return r.Send(http.MethodPost, url, bytes.NewReader(body))
}

func (r *client) Put(url string, body []byte) Response {
	r.AddHeader("Content-Type", "application/json")
	return r.Send(http.MethodPut, url, bytes.NewReader(body))
}

func (c *client) Send(method, url string, body io.Reader) Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{Error: err}
	}
	c.setHeaders(req)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return Response{Error: err}
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{Error: err}
	}
	if res.StatusCode != http.StatusOK {
		return Response{Code: res.StatusCode, Error: errors.New(c.errorMessage(content))}
	}
	return Response{
		Body:  content,
		Code:  res.StatusCode,
		Error: err}
}

func (c *client) AddHeader(key, value string) {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[key] = value
}

func (c client) errorMessage(content []byte) string {
	var err []struct {
		Codigo   string `json:"codigo"`
		Mensagem string `json:"mensagem"`
	}
	json.Unmarshal(content, &err)
	var ret []string
	for _, e := range err {
		s := fmt.Sprintf("%s %s", e.Codigo, e.Mensagem)
		ret = append(ret, s)
	}
	return strings.Join(ret, " ")
}

func (c client) Header(key string) string {
	return c.headers[key]
}

func (c *client) setAuthorization() {
	basicToken := "Basic " + c.credentials.ApiKey
	c.AddHeader("Authorization", basicToken)
}

func (c *client) setHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Accept", "application/json")
}
