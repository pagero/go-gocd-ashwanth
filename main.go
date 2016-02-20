package gocd

import "github.com/parnurzeal/gorequest"

// Client entrypoint for GoCD
type Client struct {
	Host    string `json:"host"`
	Request *gorequest.SuperAgent
}

// New GoCD Client
func New(host, username, password string) *Client {
	client := Client{
		Host:    host,
		Request: gorequest.New().SetBasicAuth(username, password),
	}
	return &client
}

func (c *Client) resolve(resource string) string {
	// TODO: Use a proper URL resolve to parse the string and append the resource
	return c.Host + resource
}