package httpHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}
type ClientOptions func(*Client) error

func New(opts ...ClientOptions) *Client {
	c := Client{}
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

func WithUrl(url string) ClientOptions {
	return func(c *Client) error {
		c.URL = url
		return nil
	}
}

func (c *Client) GetClips() ([]string, error) {
	url := fmt.Sprintf("%v/list", c.URL)
	clips := []string{}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &clips)
	if err != nil {
		return nil, err
	}
	return clips, nil
}

func (c *Client) SetClips(i int) error {
	url := fmt.Sprintf("%v/list/%v", c.URL, i)
	clip := ""
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &clip)
	if err != nil {
		return err
	}
	fmt.Println("", clip)
	return nil
}
