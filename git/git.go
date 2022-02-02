package git

import (
	"bytes"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient(opts ...ClientOption) *Client {
	return &Client{http: NewHttpClient(opts...)}
}

func (c *Client) Get(url string) string {
	if c.http == nil {
		c.http = NewHttpClient()
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}
