package e2e

import (
	"net/http"
	"testing"
)

func TestProductsEndpoint(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/products")
	if err != nil {
		t.Fatalf("failed to connect to API: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
}
