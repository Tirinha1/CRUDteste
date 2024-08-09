package crud

import (
	"github.com/Tirinha1/CRUDteste/crud/database"
	"github.com/Tirinha1/CRUDteste/crud/database/models"
)

func ReadProducts() []*models.Product {
	return database.ProductDataBase.GetAll()
}