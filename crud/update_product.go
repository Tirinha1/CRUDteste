package crud

import "github.com/Tirinha1/CRUDteste/crud/models"

func UpdateProduct(product models.Product) {
	for i, p := range models.Products {
		if p.ID == product.ID {
			models.Products[i] = product;
			break;
		}
	}
}