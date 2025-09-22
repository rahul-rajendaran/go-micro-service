package service

import "errors"

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Aspirin", Price: 10.5},
	{ID: "2", Name: "Paracetamol", Price: 8.75},
}

func GetAllProducts() []Product {
	return products
}

func CreateProduct(p Product) Product {
	p.ID = string(len(products) + 1)
	products = append(products, p)
	return p
}

func GetProductByID(id string) (Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, errors.New("not found")
}
