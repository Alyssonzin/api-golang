package httpx

import (
	"context"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	client  *http.Client
	baseURL string
	headers http.Header
}

func NewHttpClient(baseURL string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		client:  &http.Client{Timeout: timeout},
		baseURL: baseURL,
		headers: make(http.Header),
	}
}

func (c *HttpClient) SetHeader(key, value string) {
	if c.headers == nil {
		c.headers = make(http.Header)
	}
	c.headers.Set(key, value)
}

func (c *HttpClient) Do(
	ctx context.Context,
	method string,
	path string,
	headers http.Header,
	body io.Reader,
) (*http.Response, error) {
	if c.client == nil {
		c.client = &http.Client{Timeout: 10 * time.Second}
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	//mescla os headers globais do cliente com os headers específicos da requisição
	mergeHeaders(c.headers, headers)
	mergeHeaders(req.Header, headers)

	return c.client.Do(req)
}

func (c *HttpClient) Get(ctx context.Context, path string, headers http.Header) (*http.Response, error) {
	return c.Do(ctx, http.MethodGet, path, headers, nil)
}

func (c *HttpClient) Post(ctx context.Context, path string, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.Do(ctx, http.MethodPost, path, headers, body)
}

func (c *HttpClient) Put(ctx context.Context, path string, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.Do(ctx, http.MethodPut, path, headers, body)
}

func mergeHeaders(dst, src http.Header) {
	if src == nil {
		return
	}
	for k, values := range src {
		for _, v := range values {
			dst.Add(k, v)
		}
	}
}
