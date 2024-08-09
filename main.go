package main

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud"
)

func main() {
	
	crud.CreateProduct()
	crud.UpdateProduct()
	
	fmt.Println(crud.ReadProducts())
}
