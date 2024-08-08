package crud

import (
	"errors"

	"github.com/Tirinha1/CRUDteste/crud/models"
)

func SaveProduct(product models.Product) error {
	for _, p := range models.Products {
		if p.ID == product.ID {
			return errors.New("product alredy existing with this ID")
		}
	}

	models.Products = append(models.Products, product)
	return nil
}
