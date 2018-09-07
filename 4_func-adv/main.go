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

type Option func(Config) Option

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

func (c *Client) Apply(opts ...Option) []Option {
	prev := make([]Option, len(opts))
	for i := range opts {
		prev[i] = opts[i](c.c)
	}

	return prev
}

func Auth(user, pass string) Option {
	return func(c Config) Option {
		prevUser := c.user
		prevPass := c.pass
		c.user = user
		c.pass = pass

		return Auth(prevUser, prevPass)
	}
}

func Endpoint(end string) Option {
	return func(c Config) Option {
		prev := c.endpoint
		c.endpoint = end
		return Endpoint(prev)
	}
}

func Timeout(to time.Duration) Option {
	return func(c Config) Option {
		prev := c.timeout
		c.timeout = to

		return Timeout(prev)
	}
}

func DefaultTimeout(c Config) Option {
	prev := c.timeout
	c.timeout = time.Second * 30
	return Timeout(prev)
}

func main() {
	c := NewClient()
	old := c.Apply(Auth("bob", "hunter2"), DefaultTimeout, Endpoint("https://leadertech.com"))
	c.Call()
	c.Apply(old...)
	// [ api usage intensifies ]
}
