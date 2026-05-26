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

	c.client.SetHeader("X-API-KEY", apiKeyResponse.ApiKey)

	return &apiKeyResponse, nil
}

func (c *PluggyClient) GetItem(ctx context.Context, itemID string) (*dtos.GetItemResponse, error) {
	resp, err := c.client.Get(ctx, "/items/"+itemID, nil)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("pluggy status: %d, body read error: %v", resp.StatusCode, readErr)
		}
		return nil, fmt.Errorf("pluggy status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var getItemResponse dtos.GetItemResponse
	if err := json.NewDecoder(resp.Body).Decode(&getItemResponse); err != nil {
		return nil, err
	}

	return &getItemResponse, nil
}
