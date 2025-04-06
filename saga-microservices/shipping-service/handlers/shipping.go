package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shipping-service/models"
	"shipping-service/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ShipOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq struct {
		Item string `json:"item"`
	}
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shipment := models.Shipment{
		OrderID: uuid.New().String(), // ✅ Auto-generate order ID
		Item:    orderReq.Item,
		Status:  "SHIPPED",
	}
	storage.SaveShipment(shipment)
	log.Println("Order shipped:", shipment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipment)
}

func CancelShipping(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shipment.Status = "CANCELLED"
	storage.SaveShipment(shipment)
	log.Println("Shipping cancelled:", shipment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipment)
}

func GetShipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["order_id"]

	shipment, found := storage.GetShipment(orderID)
	if !found {
		http.Error(w, "Shipment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipment)
}
