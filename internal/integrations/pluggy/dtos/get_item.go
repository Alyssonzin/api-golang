package dtos

type GetItemResponse struct {
	ID              string   `json:"id"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
	Status          string   `json:"status"`
	ExecutionStatus string   `json:"executionStatus"`
	Products        []string `json:"products"`
	ClientUserID    string   `json:"clientUserId"`
}
