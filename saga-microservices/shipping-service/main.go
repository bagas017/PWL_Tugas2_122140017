package main

import (
	"log"
	"net/http"
	"shipping-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoint
	r.HandleFunc("/ship-order", handlers.ShipOrder).Methods("POST")
	r.HandleFunc("/cancel-shipping", handlers.CancelShipping).Methods("POST")
	r.HandleFunc("/shipment/{order_id}", handlers.GetShipment).Methods("GET")

	log.Println("🚚 Shipping Service running at :8002")
	if err := http.ListenAndServe(":8002", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
