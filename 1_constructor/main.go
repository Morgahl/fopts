package main

import (
	"time"
)

type Client struct {
	user     string
	pass     string
	endpoint string
	timeout  time.Duration
}

func NewClient(user, pass, endpoint string, timeout time.Duration) *Client {
	return &Client{
		user:     user,
		pass:     pass,
		endpoint: endpoint,
		timeout:  timeout,
	}
}

func (c *Client) Call() {
	// ...
}

func main() {
	c := NewClient("bob", "hunter2", "https://aviatinc.com", time.Second)
	c.Call()
	// [ api usage intensifies ]
}
