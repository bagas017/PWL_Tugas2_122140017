// main.go
package main

import (
	"log"
	"net/http"
	"orchestrator-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/start-saga", handlers.StartOrderSaga).Methods("POST")

	log.Println("Orchestrator running at :8003")
	log.Fatal(http.ListenAndServe(":8003", r))
}
