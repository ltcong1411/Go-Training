package entities

import (
	"fmt"
)

type Product struct {
	Id       string
	Name     string
	Price    float32
	Quantity int
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id: %s\nName: %s\nPrice: %f\nQuantity: %d\n", product.Id, product.Name, product.Price, product.Quantity)
}
