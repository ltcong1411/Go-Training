package main

import "fmt"

func main() {
	var a, b, sel int
	fmt.Println("Nhap vao 2 so nguyen a va b:")
	fmt.Print("nhap a:")
	fmt.Scanf("%d\n", &a)
	fmt.Print("nhap b:")
	fmt.Scanf("%d\n", &b)

	fmt.Println("a:", a, "b:", b)

	fmt.Println("Menu:\n1. a+b\n2.Dem so chan a => b\n3. Nhap vao x, kiem tra x co trong khoang a => b")
	fmt.Print("Chon yeu cau:")
	fmt.Scanf("%d\n", &sel)

	switch sel {
	case 1:
		sum(a, b)
	case 2:
		dem(a, b)
	case 3:
		check(a, b)
	default:
		fmt.Println("Why you don't choose!!!!")
	}
}

func sum(a, b int) {
	fmt.Println("a + b =", a+b)
}

func dem(a, b int) {
	sochan := 0
	for i := a; i <= b; i++ {
		fmt.Println(i)
		if (i % 2) == 0 {
			sochan++
		}
	}
	fmt.Println("Khoang a => b co", sochan, "so chan")
}

func check(a, b int) {
	fmt.Println("Nhap vao x, kiem tra x co trong khoang a => b")
}
