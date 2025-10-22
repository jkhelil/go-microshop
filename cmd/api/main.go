package main

import (
	"log"
	"net/http"

	"github.com/jkheli/go-microshop/pkg/products"
)

func main() {
	http.HandleFunc("/products", products.Handler)
	log.Println("Starting Product API on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
