package services

import (
	"fmt"
	"gestor-stock/models"
)

func MostrarStock(p models.Product) {
	fmt.Println("producto:", p.Name)
	fmt.Println("stock:", p.Stock)
}

func SumarStock(p *models.Product, cantidad int) {
	p.Stock += cantidad
}

func Listar(products []models.Product) {
	for _, p := range products {
		fmt.Printf("id: %d | Name: %s | Stock: %d\n", p.ID, p.Name, p.Stock)
		if p.Stock > 0 {
			fmt.Println("tenemos Stock")

		} else {
			fmt.Println("Sin Stock")
		}
	}
}

func FoundProduct(products []models.Product, id int) (*models.Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}
	return nil, fmt.Errorf("producto con ID %d no encontrado", id)

}

func AddProduct(products []models.Product, nuevo models.Product) []models.Product {
	return append(products, nuevo)
}
