package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/models"
	"order-service/storage"
)

func triggerPayment(order models.Order) {
	paymentReq := map[string]interface{}{
		"order_id": order.OrderID,
		"amount":   order.Amount,
	}

	body, _ := json.Marshal(paymentReq)
	resp, err := http.Post("http://localhost:8001/process-payment", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Gagal menghubungi payment service:", err)
		storage.UpdateOrderStatus(order.OrderID, "FAILED")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		storage.UpdateOrderStatus(order.OrderID, "FAILED")
	} else {
		storage.UpdateOrderStatus(order.OrderID, "PAID")
	}
}
