package crud

import (
	"github.com/Tirinha1/CRUDteste/crud/models"
)

func DeleteProduct(product models.Product) {
	for i, p := range models.Products {
		if p.ID == product.ID {
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			return
		}
	}
}