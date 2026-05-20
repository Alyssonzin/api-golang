package dtos

type CreateApiKeyRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type CreateApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}
