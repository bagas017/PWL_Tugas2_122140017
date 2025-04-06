package models

type OrderRequest struct {
	OrderID string `json:"order_id"`
	Item    string `json:"item"`
	Amount  int    `json:"amount"`
}
