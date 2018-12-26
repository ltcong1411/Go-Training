package main

import "fmt"

func main() {
	for {
		var a, b, x, sel int
		fmt.Println("Nhap vao 2 so nguyen a va b:")
		fmt.Print("nhap a:")
		fmt.Scanf("%d\n", &a)
		fmt.Print("nhap b:")
		fmt.Scanf("%d\n", &b)

		fmt.Println("a:", a, "b:", b)

		fmt.Println("Menu:\n1. a+b\n2. Dem so chan a => b\n3. Nhap vao x, kiem tra x co trong khoang a => b")
		fmt.Print("Chon yeu cau:")
		fmt.Scanf("%d\n", &sel)

		switch sel {
		case 1:
			sum(a, b)
		case 2:
			dem(a, b)
		case 3:
			fmt.Print("nhap x:")
			fmt.Scanf("%d\n", &x)
			if check(a, b, x) {
				fmt.Printf("Gia tri %d co trong khoang tu %d den %d", x, a, b)
			} else {
				fmt.Printf("Gia tri %d khong co trong khoang tu %d den %d", x, a, b)
			}
		default:
			fmt.Println("Why you don't choose!!!!")
		}
		fmt.Println("Neu muon choi tiep nhap y: ")
		var con string
		fmt.Scanf("%s\n", &con)
		if con != "y" {
			fmt.Println("Bye Bye")
			break
		}
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

func check(a, b, x int) bool {
	fmt.Println("Nhap vao x, kiem tra x co trong khoang a => b")
	for i := a; i <= b; i++ {
		if i == x {
			return true
		}
	}
	return false
}
