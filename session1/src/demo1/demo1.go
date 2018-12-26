package main

import "fmt"

func main() {
	fmt.Println("Hello")
	Demo12()
	// r1, r2, r3, r4 := calculate2(10, 4)
	// fmt.Printf("r1= %d, r2= %d, r3= %d, r4= %.2f\n", r1, r2, r3, r4)

	// display2(1, 2)
	// display2(2, 3, 4)
	// display2(2, 3, 4, 5)
}

func Demo1() {
	var age int
	var fullname string = "Le Thanh Cong"
	var status bool = true
	var price float32

	age = 20
	price = 4.5

	fmt.Println("age:", age, ", fullname:", fullname)
	fmt.Println("status:", status)
	fmt.Println("price:", price)
}

func Demo2() {
	var age = 20
	var address = "ABC"
	var a1, a2, a3 = 30, "BCD", 7

	fmt.Println("age:", age, "address:", address)
	fmt.Println("a1:", a1, "a2:", a2, "a3:", a3)
}

func Demo3() {
	age := 20
	id, name, price := "p01", "name 1", 4.5

	fmt.Println("age:", age)
	fmt.Println("id:", id, "name:", name, "price:", price)
}

func Demo4() {
	var a int = 2
	var b float32 = 4.5
	var result = float32(a) + b

	fmt.Println(result)
}

func Demo5() {
	var age int = 70
	fmt.Printf("age: %d", age)
	var status bool = true
	fmt.Printf("\nstatus: %t", status)
	var fullname string = "ABC"
	fmt.Printf("\nfullname: %s", fullname)
	var price float32 = 4.5
	fmt.Printf("\nprice: %f", price)
}

func Demo6() {
	var age int
	fmt.Print("age:")
	fmt.Scanf("%d\n", &age)
	fmt.Println("age:", age)

	var price float32
	var fullname string
	var status bool

	fmt.Print("price:")
	fmt.Scanf("%f\n", &price)
	fmt.Print("fullname:")
	fmt.Scanf("%s\n", &fullname)
	fmt.Print("status:")
	fmt.Scanf("%t\n", &status)

	fmt.Println("price:", price, "fullname:", fullname, "status:", status)
}

func Demo7() {
	var a = 5
	if a > 2 {
		fmt.Println("a > 2")
	} else {
		fmt.Println("a <= 2")
	}
}

func Demo8() {
	if a := 5; a > 2 {
		fmt.Println("a > 2")
	} else {
		fmt.Println("a <= 2")
	}
}

func Demo9() {
	menu := 1
	if menu == 1 {
		fmt.Println("Menu 1")
	} else if menu == 2 {
		fmt.Println("Menu 2")
	} else {
		fmt.Println("Invalid")
	}
}

func Demo10() {
	switch menu := 1; menu {
	case 1:
		fmt.Println("Menu 1")
	case 2:
		fmt.Println("Menu 2")
	default:
		fmt.Println("Invalid")
	}
}

func Demo11() {
	switch a := 1; a {
	case 1, 2, 3:
		fmt.Println("a = 1,2,3")
	case 4, 5, 6:
		fmt.Println("a = 4,5,6")
	default:
		fmt.Println("InValid")

	}
}

func Demo12() {
	a := 2
	switch {
	case a >= 1 && a <= 10:
		fmt.Println("a >=1 && a<=10")
	case a >= 10 && a <= 20:
		fmt.Println("a >=10 && a<=20")
	default:
		fmt.Println("InValid")
	}
}

func Demo13() {
	for i := 1; i <= 10; i++ {
		fmt.Println("  ", i)
	}
}

func Demo14_while() {
	i := 1
	for i <= 10 {
		fmt.Println("  ", i)
		i++
	}
}

func Demo15_Do_while() {
	i := 1
	for {
		fmt.Println("  ", i)
		if i > 10 {
			break
		}
		i++
	}
}

func hello(fullname string) {
	fmt.Println("Hello ", fullname)
}

func display(id, name string) {
	fmt.Println("id ", id, ", name ", name)
}

func sum(a, b int) int {
	return a + b
}

func calculate1(a, b int) (int, int, int, float32) {
	result1 := a + b
	result2 := a - b
	result3 := a * b
	result4 := float32(a) / float32(b)
	return result1, result2, result3, result4
}

func calculate2(a, b int) (result1 int, result2 int, result3 int, result4 float32) {
	result1 = a + b
	result2 = a - b
	result3 = a * b
	result4 = float32(a) / float32(b)
	return
}

func display2(args ...int) {
	fmt.Println("\nSize: ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Print("  ", args[i])
	}
}
