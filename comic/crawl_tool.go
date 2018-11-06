package comic

/*
 *设计思路为根据不同网络的json的结构
 */

import (
	"io"
	"net/http"
)

type Config struct {
	UseHTTPS bool
}

type Client struct {
	HTTPClient *http.Client
	Config
}

func NewCrawlClient() *Client {
	return &Client{
		http.DefaultClient,
		Config{
			UseHTTPS: true,
		},
	}
}

func (c *Client) rawURL(arg string) string {
	protocol := "http://"

	if c.UseHTTPS {
		protocol = "https://"
	}
	return protocol + arg
}

func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, newStatusError(res.StatusCode)
	}

	return res.Body, nil
}
