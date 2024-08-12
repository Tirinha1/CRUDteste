package main

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud"
)

func main() {

	crud.CreateProduct()

	products := crud.ReadProducts()

	crud.UpdateProduct()

	for _, product := range products {
		fmt.Println(*product)
	}
	
}
