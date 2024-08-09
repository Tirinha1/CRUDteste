package crud

import "github.com/Tirinha1/CRUDteste/crud/database"

func ReadProducts() []*database.Product {
	return database.GetAll();
}