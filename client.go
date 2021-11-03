package sample

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	logger io.Writer
	apiKey string
	ub     url.URL
	client http.Client
}

// NewClient returns client for analyze rubi
func NewClient(key string) (*Client, error) {
	client := new(Client)

	client.logger = io.Discard
	client.apiKey = key

	client.ub.Scheme = "https"
	client.ub.Host = "jlp.yahooapis.jp"
	client.ub.Path = "/FuriganaService/V2/furigana"

	return client, nil
}

func (c *Client) Close() error {
	return nil
}

// SetLogger sets logger for debug output
func (c *Client) SetLogger(logger io.Writer) {
	if logger != nil {
		c.logger = logger
	}
}

// SetAPIKey sets API key string, which is available
func (c *Client) SetAPIKey(key string) {
	c.apiKey = key
}

// SetTimeout sets timeout duration if you need
func (c *Client) SetTimeout(dur time.Duration) {
	c.client.Timeout = dur
}

// Analyze parse and
func (c *Client) Analyze(ctx context.Context, req *Request) (*Response, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	var id string
	if len(req.ID) == 0 {
		id = "ma-request"
	} else {
		id = req.ID
	}

	var buf bytes.Buffer
	err := writeRequestMessage(&buf, id, req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ub.String(), &buf)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", fmt.Sprintf("Yahoo AppID: %s", c.apiKey))

	res, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ans := new(Response)

	switch res.StatusCode {
	case http.StatusOK:
		err = readResponseMessage(res.Body, id, ans)
	case http.StatusBadRequest:
		fallthrough
	case http.StatusUnauthorized:
		fallthrough
	case http.StatusForbidden:
		fallthrough
	case http.StatusNotFound:
		fallthrough
	case http.StatusInternalServerError:
		fallthrough
	case http.StatusServiceUnavailable:
		err = readResponseMessageAsInvalid(res.Body, ans)
	default:
		err = fmt.Errorf("unhandled status code : %d", res.StatusCode)
	}

	if err != nil {
		return nil, err
	}
	return ans, nil
}
