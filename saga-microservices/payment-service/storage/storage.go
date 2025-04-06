package storage

import (
	"log"
	"payment-service/models"
)

var payments = make(map[string]models.Payment)

func SavePayment(payment models.Payment) {
	payments[payment.OrderID] = payment
	log.Printf("Saved payment: %+v\n", payment)
}

func GetPayment(orderID string) (models.Payment, bool) {
	payment, exists := payments[orderID]
	return payment, exists
}
