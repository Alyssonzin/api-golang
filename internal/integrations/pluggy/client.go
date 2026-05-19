package pluggy

import (
	httpx "api-golang/internal/httpx"
	"context"
	"fmt"
	"io"
	"time"
)

type PluggyClient struct {
	client *httpx.HttpClient
}

func NewPluggyClient(baseURL string) *PluggyClient {
	c := httpx.NewHttpClient(baseURL, 5*time.Second)
	c.SetHeader("Accept", "application/json")

	return &PluggyClient{
		client: c,
	}
}

func (c *PluggyClient) GetItem(ctx context.Context) ([]byte, error) {
	resp, err := c.client.Get(ctx, "/healthz", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("pluggy status: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
