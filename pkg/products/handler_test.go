package products

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ReturnsProducts(t *testing.T) {
	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	Handler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", res.StatusCode)
	}

	var items []Product
	if err := json.NewDecoder(res.Body).Decode(&items); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(items) == 0 {
		t.Errorf("Expected at least 1 product, got 0")
	}
}
