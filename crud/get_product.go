package crud

import "github.com/Tirinha1/CRUDteste/crud/models"

// Retorna todos os produtos salvos
func GetProducts() []models.Product {
	return models.Products
}