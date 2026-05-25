package usecase

import (
	"api-golang/internal/integrations/pluggy"
	"context"
)

type PluggyApikeyUseCase struct {
	pluggyClient *pluggy.PluggyClient
}

func NewPluggyApikeyUseCase(pluggyClient *pluggy.PluggyClient) PluggyApikeyUseCase {
	return PluggyApikeyUseCase{pluggyClient: pluggyClient}
}

func (uc *PluggyApikeyUseCase) CreatePluggyApikey(ctx context.Context) (string, error) {
	res, err := uc.pluggyClient.CreateApiKey(ctx)
	if err != nil {
		return "", err
	}
	return res.ApiKey, nil
}
