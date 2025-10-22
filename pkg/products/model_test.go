package products

import (
	"encoding/json"
	"testing"
)

func TestProductJSONMarshalling(t *testing.T) {
	product := Product{ID: 1, Name: "Keyboard", Price: 49.99}

	data, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("Failed to marshal product: %v", err)
	}

	var decoded Product
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Name != product.Name {
		t.Errorf("Expected %s, got %s", product.Name, decoded.Name)
	}
}
