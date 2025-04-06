package models

type Shipment struct {
	OrderID string `json:"order_id"`
	Item    string `json:"item"`
	Status  string `json:"status"`
}
