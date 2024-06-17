package main

import (
	"encoding/json"
	"net/http"
)

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// buat route
	mux := http.NewServeMux()

	// Handler
	// tambahkan handlre ke mux

	mux.HandleFunc("GET /products", listProductFunc)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// Membuat Server
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	// MenjalanKan Server
	server.ListenAndServe()
}

var database = map[int]Products{}

func listProductFunc(w http.ResponseWriter, r *http.Request) {
	var products []Products

	for _, v := range database {
		products = append(products, v)
	}
	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
}
