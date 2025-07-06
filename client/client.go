package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://jsonplaceholder.typicode.com/"
)

type Post struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}
type HTTPClient struct {
	client  *http.Client
	BaseURL *url.URL
}

func NewHTTPClient(baseClient *http.Client) *HTTPClient {
	if baseClient == nil {
		baseClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	return &HTTPClient{
		client:  baseClient,
		BaseURL: baseURL,
	}
}

func (c *HTTPClient) NewRequest(method, urlStr string, body any) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %v does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *HTTPClient) Do(ctx context.Context, req *http.Request, v any) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}

	return resp, err
}

func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return fmt.Errorf("%v %v : %v", resp.Request.Method, resp.Request.URL, resp.Status)
}

func (c *HTTPClient) GetPost(ctx context.Context, id int64) (*Post, *http.Response, error) {
	u := fmt.Sprintf("posts/%v", id)

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	structPost := new(Post)
	resp, err := c.Do(ctx, req, structPost)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	return structPost, resp, nil

}
