package client

import (
	"encoding/json"
	"fmt"
	"github.com/dan-ibm/go-client-beer/response"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	BaseURL = "https://rustybeer.herokuapp.com"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() (*Client, error) {
	return &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) sendRequest(req *http.Request, data interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code %d\n", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return nil
}

func (c *Client) GetStyles() ([]response.Beer, error) {
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/styles", c.BaseURL))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r []response.Beer

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetStylesByName(name string) ([]response.Beer, error) {
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/styles?name=%s", c.BaseURL, name))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r []response.Beer

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}
