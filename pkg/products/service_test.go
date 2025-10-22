package products

import "testing"

func TestGetProductByID(t *testing.T) {
	p, err := GetProductByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.Name != "Laptop" {
		t.Errorf("expected 'Laptop', got '%s'", p.Name)
	}

	_, err = GetProductByID(99)
	if err == nil {
		t.Error("expected error for missing product, got nil")
	}
}
