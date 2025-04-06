package models

type Payment struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"` // SUCCESS, FAILED, REFUNDED
}
