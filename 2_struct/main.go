package main

import (
	"time"
)

type Client struct {
	c Config
}

type Config struct {
	User     string
	Pass     string
	Endpoint string
	Timeout  time.Duration
}

func NewClient(conf Config) *Client {
	return &Client{
		c: conf,
	}
}

func (c *Client) Call() {
	// ...
}

func main() {
	conf := Config{
		User:     "bob",
		Pass:     "hunter2",
		Endpoint: "https://aviatinc.com",
		Timeout:  time.Second,
	}
	c := NewClient(conf)
	c.Call()
	// [ api usage intensifies ]
}
