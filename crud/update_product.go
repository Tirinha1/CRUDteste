package crud

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud/models"
)

func UpdateProduct() {
	
	var id int

	fmt.Println("Qual o ID do produto que deseja alterar?: ")
	fmt.Scanf("%d", &id)

	product := models.GetProduct(id)

	var opc int 
	fmt.Println("O que deseja alterar do produto: 1-Preço, 2-Pais, 3-Nome")
	fmt.Scanf("%d", &opc)

	switch opc {

	case 1:
		fmt.Println("Insira o novo preço: ")
		fmt.Scanf("%d", &product.Price)

		models.UpdtProduct(product);
	case 2:
		fmt.Println("Insira o novo pais: ")
		fmt.Scanf("%v", &product.ManufactureCountry)

		models.UpdtProduct(product);
	case 3:
		fmt.Println("Insira o novo nome: ")
		fmt.Scanf("%v", &product.Name)

		models.UpdtProduct(product);
	default:
		fmt.Println("Digite uma opção valida!")
	}
}