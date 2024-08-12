package crud

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud/database"
)

func UpdateProduct() {

	var id int

	// TODO: ReadProducts()
	products := ReadProducts()
	for _, product := range products {
		fmt.Println(*product)
	}

	fmt.Println("Qual o ID do produto que deseja alterar?: ")
	fmt.Scanf("%d\n", &id)
	product, err := database.ProductDataBase.Get(id)
	if err != nil {
		fmt.Print("erro")
	} else {
		opc := 5
		for {
			if opc != 0 {
				fmt.Println("Escolha uma opção:")
				fmt.Println("1. Atualizar preço")
				fmt.Println("2. Atualizar país de fabricação")
				fmt.Println("3. Atualizar nome")
				fmt.Println("0. Sair")
				fmt.Scanf("%d\n", &opc)

				switch opc {
				case 1:
					fmt.Println("Insira o novo preço: ")
					fmt.Scanf("%v", &product.Price)
					database.ProductDataBase.Update(product)
					return
				case 2:
					fmt.Println("Insira o novo país: ")
					fmt.Scanf("%v", &product.ManufactureCountry)
					database.ProductDataBase.Update(product)
					return
				case 3:
					fmt.Println("Insira o novo nome: ")
					fmt.Scanf("%v", &product.Name)
					database.ProductDataBase.Update(product)
					return

				default:
					fmt.Println("Digite uma opção válida!")
				}
			}
		}
	}
}
