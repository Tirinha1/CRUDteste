package main

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud"
	"github.com/Tirinha1/CRUDteste/crud/models"
)

func main() {
	var firstProduct models.Product
	crud.CreateProduct(firstProduct)

	var secondProduct models.Product
	crud.CreateProduct(secondProduct)


	fmt.Println("The current procuts are: ", crud.ReadProducts())

}
