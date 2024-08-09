package crud

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud/database"
	"github.com/Tirinha1/CRUDteste/crud/database/models"
)

func CreateProduct() {

	product := models.Product{}

	fmt.Printf("Digite o nome do produto: ")
	fmt.Scanf("%v\n", &product.Name)

	fmt.Printf("Digite o preço do produto: ")
	fmt.Scanf("%v\n", &product.Price)

	fmt.Printf("Digite o pais que produz o produto: ")
	fmt.Scanf("%v\n", &product.ManufactureCountry)

	database.ProductDataBase.Add(&product)
}
