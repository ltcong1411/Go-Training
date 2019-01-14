package entities

<<<<<<< HEAD
import "fmt"

type Product struct {
	Id, Name string
=======
import (
	"fmt"
)

type Product struct {
	Id       string
	Name     string
>>>>>>> b16df0e3461b618480359a36a8992fd0fd167abc
	Price    float32
	Quantity int
}

func (product Product) ToString() string {
<<<<<<< HEAD
	return fmt.Printf("id: %s\nname: %s\nprice: %f\nquantity: %d", product.Id, product.Name, product.Price, product.Quantity)
}

func (product Product) Total() float32 {
	return product.Price * float32(product.Quantity)
=======
	return fmt.Sprintf("Id: %s\nName: %s\nPrice: %f\nQuantity: %d\n", product.Id, product.Name, product.Price, product.Quantity)
>>>>>>> b16df0e3461b618480359a36a8992fd0fd167abc
}
