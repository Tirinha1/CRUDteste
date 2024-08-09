package crud

import "github.com/Tirinha1/CRUDteste/crud/models"

func ReadProducts() []*models.Product {
	return models.GetAll();
}