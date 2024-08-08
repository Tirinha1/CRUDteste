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

	firstProduct.Price = 350;

	crud.UpdateProduct(firstProduct);
	crud.DeleteProduct(firstProduct);

	fmt.Println("Current products in memory:", crud.ReadProducts())
}
