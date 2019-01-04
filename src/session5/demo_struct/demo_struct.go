package demo_struct

import (
	"fmt"
	"session5/entities"
)

func Demo1() {
	var student entities.Student
	student.Id = "st01"
	student.Name = "Name 1"
	student.Age = 20
	student.Score = 6
	fmt.Println(student)
	fmt.Println("Id:", student.Id)
}

func Demo2() {
	student := entities.Student{
		Id:    "st02",
		Name:  "Name 2",
		Age:   20,
		Score: 6,
	}
	fmt.Println(student)
}

func Demo3() {
	student := entities.Student{"st03", "Name 3", 23, 10}
	fmt.Println(student)
}

func Demo4() {
	var employee entities.Employee
	employee.Id = "e01"
	employee.FullName = entities.FullName{
		FirstName:  "FN 1",
		MiddleName: "MD 1",
		LastName:   "LN 1",
	}

	employee.Address = entities.Address{
		Street:   "Street 1",
		Ward:     "Ward 1",
		District: "Dis 1",
	}

	employee.Salary = 20
	fmt.Println(employee)
	fmt.Println("First Name:", employee.FullName.FirstName)
}

func Demo6() {
	student := entities.Student{"st03", "Name 3", 23, 10}
	var p *entities.Student = &student
	fmt.Println("Name: ", (*p).Name)
	(*p).Age = 20
	fmt.Println(student)
}

func Change(p *entities.Student) {
	(*p).Name = "abc"
}

func Demo7() {
	student := entities.Student{"st03", "Name 3", 23, 10}
	Change(&student)
	fmt.Println(student)
}

func Demo8() {
	students := [3]entities.Student{
		{"st01", "name 1", 20, 7},
		{"st02", "name 1", 20, 7},
		{"st03", "name 1", 20, 7},
	}
	for _, student := range students {
		fmt.Println(ToString(student))
		fmt.Println("...........................")
	}
}

func ToString(student entities.Student) string {
	return fmt.Sprintf("id: %s\nname: %s\nage:%d\nscore:%d", student.Id, student.Name, student.Age, student.Score)
}

/*
Khai bao struct Product co cac thuoc tinh sau:
1. Id kieu chuoi
2. Name kieu chuoi
3. Price kieu so thuc
4. Quantity kieu so nguyen
Khai bao mang chua 5 sp va khai bao cac ham thuc hien cac hanh dong sau:
a. Tinh tong tien tat ca cac sp
b. tim sp cos gia max, min
c. Sap xep cac sp theo gia tang dan, giam dan
d. Tim cac sp co ten chua 1 keyword
*/
func Demo() {
	products := [5]entities.Product{
		{"p01", "Apple", 10, 1},
		{"p02", "Samsung", 20, 2},
		{"p03", "Sony", 30, 3},
		{"p03", "Xiaomi", 40, 4},
		{"p03", "HTC", 50, 5},
	}

	for _, product := range products {
		fmt.Println(ToStringProduct(product))
		fmt.Println("...........................")
	}

}

func ToStringProduct(product entities.Product) string {
	return fmt.Sprintf("id: %s\nname: %s\nprice:%d\nquantity:%d", product.Id, product.Name, product.Price, product.Quantity)
}

// func Question1(products []entities.Product) total int {

// }
