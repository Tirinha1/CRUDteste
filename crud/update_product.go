package crud

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud/database"
)

func UpdateProduct() {
	
	var id int

	// TODO: ReadProducts()

	fmt.Println("Qual o ID do produto que deseja alterar?: ")
	fmt.Scanf("%d", &id)

	product, err := database.ProductDataBase.Get(id)
	if err != nil {
		fmt.Print("erro")
	}

	var opc int 
	fmt.Println("O que deseja alterar do produto: 1-Preço, 2-Pais, 3-Nome")
	fmt.Scanf("%d", &opc)

	switch opc {

	case 1:
		fmt.Println("Insira o novo preço: ")
		fmt.Scanf("%d", &product.Price)

		database.ProductDataBase.Update(product);
	case 2:
		fmt.Println("Insira o novo pais: ")
		fmt.Scanf("%v", &product.ManufactureCountry)

		database.ProductDataBase.Update(product);
	case 3:
		fmt.Println("Insira o novo nome: ")
		fmt.Scanf("%v", &product.Name)

		database.ProductDataBase.Update(product);
	default:
		fmt.Println("Digite uma opção valida!")
	}
}