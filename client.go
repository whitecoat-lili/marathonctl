package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Login struct {
	Host string // todo hosts
	User string
	Pass string
}

func NewLogin(host, login string) *Login {
	toks := strings.SplitN(login, ":", 2)
	return &Login{
		Host: host,
		User: toks[0],
		Pass: toks[1],
	}
}

func Usage() {
	fmt.Println(Howto)
	os.Exit(1)
}

type Tool struct {
	actions map[string]Action
}

func (t *Tool) Start(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if action, ok := t.actions[args[0]]; !ok {
		Usage()
	} else {
		action.Apply(args[1:])
	}
}

type Client struct {
	client http.Client
	login  *Login
}

func NewClient(login *Login) *Client {
	return &Client{
		client: http.Client{},
		login:  login,
	}
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	return c.client.Do(r)
}

func (c *Client) GET(path string) *http.Request {
	url := c.login.Host + path
	r, e := http.NewRequest("GET", url, nil)
	Check(e == nil, "failed to crete get request", e)
	r.SetBasicAuth(c.login.User, c.login.Pass)
	return r
}

func (c *Client) POST(path string, body io.ReadCloser) *http.Request {
	url := c.login.Host + path
	r, e := http.NewRequest("POST", url, body)
	Check(e == nil, "failed to create post request", e)
	r.Header.Set("Content-Type", "application/json")
	r.SetBasicAuth(c.login.User, c.login.Pass)
	return r
}

func (c *Client) DELETE(path string) *http.Request {
	url := c.login.Host + path
	r, e := http.NewRequest("DELETE", url, nil)
	Check(e == nil, "failed to create delete request", e)
	r.Header.Set("Content-Type", "application/json")
	r.SetBasicAuth(c.login.User, c.login.Pass)
	return r
}
