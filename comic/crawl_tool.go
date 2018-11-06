package comic

/*
 *设计思路为根据不同网络的json的结构
 */

import (
	"io"
	"net/http"
)

type Config struct {
	UseHttps bool
}

type Client struct {
	HttpClient *http.Client
	Config
}

func NewCrawlClient() *Client {
	return &Client{
		http.DefaultClient,
		Config{
			UseHttps: true,
		},
	}
}

func (c *Client) rawURL(arg string) string {
	protocol := "http://"

	if c.UseHttps {
		protocol = "https://"
	}
	return protocol + arg
}

func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, newStatusError(res.StatusCode)
	}

	return res.Body, nil
}
