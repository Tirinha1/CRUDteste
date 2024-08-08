package main

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud"
	"github.com/Tirinha1/CRUDteste/crud/models"
)
var firstProduct models.Product
var secondProduct models.Product

func main() {
	crud.CreateProduct(firstProduct, 1);
	crud.CreateProduct(secondProduct, 2);
	
	fmt.Println("The current procuts are: ", crud.ReadProducts())

	crud.DeleteProduct(models.Products[0]);

}
