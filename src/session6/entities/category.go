package entities

type Category struct {
	Id, Name string
	Products []Product
}

func (category Category) ToString() (result string) {
	result = "Id: " + category.Id
	result += "\nName: " + category.Name
	result += "\nProducts: \n"
	for _, product := range category.Products {
		result += product.ToString()
		result += "------------\n"
	}
	return
}

func (category Category) ToTal() (result float32) {
	result = 0
	for _, product := range category.Products {
		result += product.Price
	}
	return
}
