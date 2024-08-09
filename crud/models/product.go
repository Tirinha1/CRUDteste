package models

import "math/rand/v2"

type Product struct {
	ID                 int
	Name               string
	ManufactureCountry string
	Price              int
}

var Products []Product

var products2 []*Product = make([]*Product, 0)

func AddProduct(product *Product) {
	id := rand.IntN(9999999)
	product.ID = id
	if GetProduct(id) == nil {
		products2 = append(products2, product)
	}
}

func DelProduct(id int) {
	for i, p := range products2 {
		if p.ID == id {
			products2 = append(products2[:i], products2[i+1:]...)
		}
	}
}

func UpdtProduct(product *Product) {
	for i, p := range products2 {
		if p.ID == product.ID {
			products2[i] = product
		}
	}
}

func GetProduct(id int) *Product {
	for _, p := range products2 {
		if p.ID == id {

			return p
		}
	}
	return nil
}

func GetAll() []*Product {
	return products2
}
