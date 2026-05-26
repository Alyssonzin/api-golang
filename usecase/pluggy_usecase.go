package usecase

import (
	"api-golang/internal/integrations/pluggy"
	"api-golang/internal/integrations/pluggy/dtos"
	"context"
	"fmt"
)

type PluggyUseCase struct {
	pluggyClient *pluggy.PluggyClient
}

func NewPluggyUseCase(pluggyClient *pluggy.PluggyClient) PluggyUseCase {
	return PluggyUseCase{pluggyClient: pluggyClient}
}

func (uc *PluggyUseCase) CreatePluggyApikey(ctx context.Context) (string, error) {
	res, err := uc.pluggyClient.CreateApiKey(ctx)
	if err != nil {
		return "", err
	}
	return res.ApiKey, nil
}

func (uc *PluggyUseCase) GetItem(ctx context.Context, itemID string) (*dtos.GetItemResponse, error) {
	res, err := uc.pluggyClient.GetItem(ctx, itemID)
	if err != nil {
		fmt.Printf("Error getting item: %v\n", err)
		return nil, err
	}
	return res, nil
}
