package main

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud"
	"github.com/Tirinha1/CRUDteste/crud/models"
)

func main() {

	firstProduct := models.Product {
		ID: 1,
		Name: "Smartphone",
		ManufactureCountry: "Brazil",
		Price: 299,
	}

	crud.SaveProduct(firstProduct);

	secondProduct := models.Product {
		ID: 2,
		Name: "Smartphone",
		ManufactureCountry: "Brazil",
		Price: 299,
	}

	crud.SaveProduct(secondProduct);

	fmt.Println("Current products in memory:", crud.GetProducts())
}
