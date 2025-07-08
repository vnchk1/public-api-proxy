package client

import (
	"bytes"
	"context"
	"encoding/json"
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

func NewHTTPClient(baseClient *http.Client) (*HTTPClient, error) {
	if baseClient == nil {
		return nil, fmt.Errorf("baseClient is nil")
	}
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL error: %v", err)
	}
	if !strings.HasSuffix(baseURL.String(), "/") {
		return nil, fmt.Errorf("URL must have a trailing slash %v", baseURL.String())
	}
	return &HTTPClient{
		client:  baseClient,
		BaseURL: baseURL,
	}, nil
}

func (c *HTTPClient) NewRequest(method, urlStr string, body any) (*http.Request, error) {

	reqUrl, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("error encoding request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, reqUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *HTTPClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	err = CheckResponse(resp)

	return resp, err
}

func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; http.StatusOK <= c && c <= http.StatusIMUsed {
		return nil
	}

	return fmt.Errorf("%v %v : %v", resp.Request.Method, resp.Request.URL, resp.Status)
}

func (c *HTTPClient) GetPost(ctx context.Context, id int64) (*Post, *http.Response, error) {
	idEndpoint := fmt.Sprintf("posts/%v", id)

	req, err := c.NewRequest(http.MethodGet, idEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	structPost := new(Post)
	resp, err := c.Do(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(structPost)
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, nil, err
	}
	return structPost, resp, nil

}
