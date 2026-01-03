package main

import (
	"fmt"
	"gestor-stock/models"
	"gestor-stock/services"
)

func main() {

	products := []models.Product{
		{ID: 2, Name: "fideos", Stock: 10},
		{ID: 1, Name: "arroz", Stock: 20},
	}

	services.Listar(products)

	p, err := services.FoundProduct(products, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.Name)

	products = services.AddProduct(products, models.Product{
		ID: 3, Name: "azucar", Stock: 50})

	services.Listar(products)

}

// var name string = "david"
// var age int = 35
// weight := 95.5

// fmt.Println("mi name is:", name, "i have:", age, "years old", "my weight is:", weight)

// if age == 35 {
// 	fmt.Println("you're so old")
// } else {
// 	fmt.Println("you're young")
// }
