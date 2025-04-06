package storage

import "order-service/models"

var Orders = make(map[string]models.Order)

func SaveOrder(order models.Order) {
	Orders[order.OrderID] = order
}

func GetOrder(orderID string) (models.Order, bool) {
	order, exists := Orders[orderID]
	return order, exists
}

func UpdateOrderStatus(orderID, status string) {
	if order, exists := Orders[orderID]; exists {
		order.Status = status
		Orders[orderID] = order
	}
}
