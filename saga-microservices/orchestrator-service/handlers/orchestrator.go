package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"orchestrator-service/models"
)

func StartOrderSaga(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1. Buat Order
	orderRes, err := http.Post("http://localhost:8000/create-order", "application/json", toJson(req))
	if err != nil || orderRes.StatusCode != http.StatusOK {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// 2. Proses Pembayaran
	paymentRes, err := http.Post("http://localhost:8001/process-payment", "application/json", toJson(req))
	if err != nil || paymentRes.StatusCode != http.StatusOK {
		// kompensasi: cancel order
		http.Post("http://localhost:8000/cancel-order", "application/json", toJson(req))
		http.Error(w, "Failed to process payment", http.StatusInternalServerError)
		return
	}

	// 3. Kirim Barang (pakai struct Shipment agar ada field Item)
	shipment := struct {
		OrderID string `json:"order_id"`
		Item    string `json:"item"`
	}{
		OrderID: req.OrderID,
		Item:    req.Item,
	}
	shippingRes, err := http.Post("http://localhost:8002/ship", "application/json", toJson(shipment))
	if err != nil || shippingRes.StatusCode != http.StatusOK {
		// kompensasi berurutan
		http.Post("http://localhost:8002/cancel-shipping", "application/json", toJson(shipment))
		http.Post("http://localhost:8001/refund-payment", "application/json", toJson(req))
		http.Post("http://localhost:8000/cancel-order", "application/json", toJson(req))
		http.Error(w, "Failed to ship item", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order processed successfully!"))
}

func toJson(data interface{}) *bytes.Buffer {
	b, _ := json.Marshal(data)
	return bytes.NewBuffer(b)
}
