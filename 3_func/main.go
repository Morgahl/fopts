package main

import (
	"time"
)

type Client struct {
	c Config
}

type Config struct {
	user     string
	pass     string
	endpoint string
	timeout  time.Duration
}

type Option func(Config)

func NewClient(opts ...Option) *Client {
	return &Client{
		c: Config{
			endpoint: "https://aviatinc.com",
			timeout:  time.Second * 30,
		},
	}
}

func (c *Client) Call() {
	// ...
}

func (c *Client) Apply(opts ...func(Config)) {
	for i := range opts {
		opts[i](c.c)
	}
}

func Auth(user, pass string) func(Config) {
	return func(c Config) {
		c.user = user
		c.pass = pass
	}
}

func Endpoint(end string) func(Config) {
	return func(c Config) {
		c.endpoint = end
	}
}

func Timeout(to time.Duration) func(Config) {
	return func(c Config) {
		c.timeout = to
	}
}

func DefaultTimeout(c Config) {
	c.timeout = time.Second * 30
}

func main() {
	c := NewClient()
	c.Apply(Auth("bob", "hunter2"), DefaultTimeout, Endpoint("https://leadertech.com"))
	c.Call()
	// [ api usage intensifies ]
}
