package main

import (
	"log"
	"net/http"
	"order-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-order", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/cancel-order", handlers.CancelOrderHandler).Methods("POST")
	r.HandleFunc("/get-order", handlers.GetOrderHandler).Methods("GET")

	log.Println("Order Service running on port 8000")
	http.ListenAndServe(":8000", r)
}
