package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/models"
	"order-service/storage"

	"github.com/google/uuid"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	orderID := uuid.New().String()
	order := models.Order{
		OrderID: orderID,
		Amount:  req.Amount,
		Status:  "PENDING",
	}

	storage.SaveOrder(order)

	paymentReq := map[string]interface{}{
		"order_id": orderID,
		"amount":   req.Amount,
	}

	jsonData, err := json.Marshal(paymentReq)
	if err != nil {
		http.Error(w, "Failed to create payment request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:8001/process-payment", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Gagal menghubungi Payment Service:", err)
		storage.UpdateOrderStatus(orderID, "FAILED")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Payment gagal diproses oleh Payment Service")
		storage.UpdateOrderStatus(orderID, "FAILED")
	} else {
		storage.UpdateOrderStatus(orderID, "PAID")
	}

	order, _ = storage.GetOrder(orderID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID string `json:"order_id"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	storage.UpdateOrderStatus(req.OrderID, "CANCELLED")
	w.WriteHeader(http.StatusOK)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	order, found := storage.GetOrder(orderID)
	if !found {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
