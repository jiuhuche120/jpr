package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jiuhuche120/jpr/config"
)

type ClientOption = func(*Client)
type Client struct {
	http   *http.Client
	Header map[string]string
}

func NewClient(opts ...ClientOption) *Client {
	client := &Client{http: http.DefaultClient, Header: map[string]string{}}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func AddHeader(key, value string) ClientOption {
	return func(client *Client) {
		client.Header[key] = value
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	if c.http == nil {
		c.http = http.DefaultClient
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}

func (c *Client) Post(url string, data interface{}) ([]byte, error) {
	if c.http == nil {
		c.http = http.DefaultClient
	}
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}

func (c *Client) GetPullRequests(url string) ([]PullRequest, error) {
	data, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	var pulls []PullRequest
	err = json.Unmarshal(data, &pulls)
	if err != nil {
		return nil, err
	}
	return pulls, nil
}

func (c *Client) GetAllPullRequests(gits config.Gits) ([]PullRequest, error) {
	return c.GetPullRequestByStatus(gits, "all")
}

func (c *Client) GetPullRequestByStatus(gits config.Gits, status string) ([]PullRequest, error) {
	return c.GetPullRequests(getUrl(gits.Owner, gits.Repo) + "?state=" + status)
}

func getUrl(owner, repo string) string {
	return "https://api.github.com/repos/" + owner + "/" + repo + "/pulls"
}
