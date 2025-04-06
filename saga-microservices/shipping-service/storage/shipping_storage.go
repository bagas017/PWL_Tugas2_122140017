package storage

import (
	"log"
	"shipping-service/models"
)

var shipmentData = make(map[string]models.Shipment)

func SaveShipment(s models.Shipment) {
	shipmentData[s.OrderID] = s
	log.Println("Saved shipment:", s)
}

func GetShipment(orderID string) (models.Shipment, bool) {
	s, ok := shipmentData[orderID]
	return s, ok
}
