package pluggy

import (
	httpx "api-golang/internal/httpx"
	dtos "api-golang/internal/integrations/pluggy/dtos"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type PluggyClient struct {
	client       *httpx.HttpClient
	clientID     string
	clientSecret string
}

func NewPluggyClient(clientID, clientSecret string) *PluggyClient {
	c := httpx.NewHttpClient("https://api.pluggy.ai", 5*time.Second)
	c.SetHeader("Accept", "application/json")

	return &PluggyClient{
		client:       c,
		clientID:     clientID,
		clientSecret: clientSecret,
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

func (c *PluggyClient) CreateApiKey(ctx context.Context) (*dtos.CreateApiKeyResponse, error) {
	req := dtos.CreateApiKeyRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(payload)

	resp, err := c.client.Post(ctx, "/auth", body, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("pluggy status: %d", resp.StatusCode)
	}

	var apiKeyResponse dtos.CreateApiKeyResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiKeyResponse); err != nil {
		return nil, err
	}

	return &apiKeyResponse, nil
}
