package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func getId(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/api/endpoint", getData).Methods("GET")
	r.HandleFunc("/api/endpoint/{id}", getId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
