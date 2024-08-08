package crud

import (
	"github.com/Tirinha1/CRUDteste/crud/models"
)

func CreateProduct(product models.Product, amount int) {
	product.ID = amount
	product.Name = "Smartphone"
    product.ManufactureCountry = "Brazil"
	product.Price = 399

    SaveProduct(product)
}