// internal/interface/http/dto/order_dto.go

package dto

import "time"

type CreateOrderRequest struct {
	ItemID   string  `json:"item_id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type CreateOrderResponse struct {
	ID        string    `json:"id"`
	ItemID    string    `json:"item_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type GetOrderResponse struct {
	ID        string    `json:"id"`
	ItemID    string    `json:"item_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
