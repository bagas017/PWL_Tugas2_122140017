package main

import (
	"log"
	"net/http"
	"payment-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoint untuk memproses pembayaran
	r.HandleFunc("/process-payment", handlers.ProcessPayment).Methods("POST")

	// Endpoint untuk refund
	r.HandleFunc("/refund-payment", handlers.RefundPayment).Methods("POST")

	log.Println("Payment Service running at :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
