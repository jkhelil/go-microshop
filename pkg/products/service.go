package products

import "errors"

var productsDB = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Headphones", Price: 199.99},
}

func GetProductByID(id int) (*Product, error) {
	for _, p := range productsDB {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}
