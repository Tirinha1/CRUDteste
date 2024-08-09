package database

import "github.com/Tirinha1/CRUDteste/crud/database/models"

// T = *models.Product
var ProductDataBase = Database[*models.Product]{}
// var PersonDatabase = Database[*models.Person]{} <<<<< 
// product.SetId()