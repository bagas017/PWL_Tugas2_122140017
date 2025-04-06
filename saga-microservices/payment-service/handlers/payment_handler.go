package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"payment-service/models"
	"payment-service/storage"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment.Status = "SUCCESS" // Simulasi sukses
	storage.SavePayment(payment)
	log.Println("Payment processed:", payment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func RefundPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment.Status = "REFUNDED"
	storage.SavePayment(payment)
	log.Println("Payment refunded:", payment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}
