package crud

import (
	"fmt"

	"github.com/Tirinha1/CRUDteste/crud/models"
)

func DeleteProduct() {
	var id int

	fmt.Println("Qual o ID do produto que deseja deletar?: ")
	fmt.Scanf("%d", &id)

	models.DelProduct(id)
}
