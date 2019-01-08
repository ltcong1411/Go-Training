package main

import (
	"fmt"
	"session6/entities"
)

func main() {
	Demo8()
}

func Demo1() {
	student1 := entities.Student{"st01", "Name 1", 20}
	fmt.Println(student1.ToString())
}

func Demo2() {
	product1 := entities.Product{"p01", "Name 1", 2, 5}
	fmt.Println(product1.ToString())
}

func Demo3() {
	student1 := entities.NewStudent("st01", "Name 1", 20)
	fmt.Println(student1.ToString())
}

func Demo4() {
	student1 := entities.NewStudent("st01", "Name 1", 20)
	fmt.Println(student1.ToString())
	student1.SetName("abc")
	fmt.Println(student1.ToString())
	(&student1).SetAge(13)
	fmt.Println(student1.ToString())
}

func Demo5() {
	category1 := entities.Category{"c1", "name 1", []entities.Product{
		{"p01", "Name 1", 2, 5},
		{"p02", "Name 2", 2, 5},
	}}

	fmt.Println(category1.ToString())
}

/*
Tao mang category co 2 danh muc. Xay dung ham thuc hien cac yeu cau sau:
1. Tinh tong tien tat ca san pham trong cac danh muc
2. Tim thong tin sp lon nhat trong mot danh muc
3. Tim thong tin sp lon nhat trong tat ca cac danh muc
*/

// func Demo() {
// 	categories := []entities.Category{
// 		{"c1", "name 1", []entities.Product{
// 			{"p01", "Name 1", 2, 5},
// 			{"p02", "Name 2", 2, 10},
// 		}},
// 		{"c2", "name 2", []entities.Product{
// 			{"p03", "Name 3", 2, 3},
// 			{"p04", "Name 4", 2, 7},
// 		}},
// 	}
// 	fmt.Println("Question1: ", Question1(categories))
// }

// func Question1(categories []Category) float32 {
// 	return 2
// }

func Demo6() {
	employee := entities.Employee{entities.Human{"name 1", "male"}, "e01", 20}
	fmt.Println(employee.ToString())
}

func Demo7() {
	var human entities.Human
	var Ihuman1 entities.IHuman1 = human
	Ihuman1.Sleep()
	fmt.Println(Ihuman1.Eat())
	var Ihuman2 entities.IHuman2 = human
	Ihuman2.Move()
}

func Demo8() {
	square1 := entities.Square{"Square1", 5}
	rectagle1 := entities.Rectagle{"Rectagle1", 5, 10}
	geomatries := []entities.Geomatry{square1, rectagle1}
	for _, geomatry := range geomatries {
		fmt.Println(geomatry.Type())
		fmt.Println("Area: ", geomatry.Area())
		fmt.Println("Perimeter: ", geomatry.Perimeter())
		fmt.Println("---------------")
	}
}
