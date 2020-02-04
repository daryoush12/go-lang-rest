package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server")
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"Server Status": true})
}
