package client

import (
	"time"
)

type Client struct {
	client   Caller
	Endpoint string
	timeout  time.Duration
	debug    bool
	key      string
}

func New(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
		timeout:  10 * time.Second,
		client:   DefaultFastHTTPCaller,
	}
}

// timeout for requests
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.timeout = timeout
	return c
}

// print debug info
func (c *Client) WithDebug() *Client {
	c.debug = true
	c.client.WithDebug()
	return c
}

// update endpoint
func (c *Client) WithEndpoint(endpoint string) *Client {
	c.Endpoint = endpoint
	return c
}

func (c *Client) WithAPIKey(key string) *Client {
	c.key = key
	return c
}
