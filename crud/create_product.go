package crud

import (
	"fmt"
	"github.com/Tirinha1/CRUDteste/crud/database"
)

func CreateProduct() {

	product := database.Product{}

	fmt.Printf("Digite o nome do produto: ")
	fmt.Scanf("%v\n", &product.Name)

	fmt.Printf("Digite o preço do produto: ")
	fmt.Scanf("%v\n", &product.Price)

	fmt.Printf("Digite o pais que produz o produto: ")
	fmt.Scanf("%v\n", &product.ManufactureCountry)

	database.AddProduct(&product)
}
