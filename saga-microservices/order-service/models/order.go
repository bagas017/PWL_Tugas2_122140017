package models

type Order struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"` // PENDING, COMPLETED, CANCELLED
}
