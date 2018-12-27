package demo_pointer

import "fmt"

func Demo1() {
	var a int = 5
	fmt.Println(5)
	fmt.Println(&a)
}

func Demo2() {
	var a int = 5
	var p *int = &a //pointer

	//address
	fmt.Println(&a)
	fmt.Println(p)

	//value in this address
	fmt.Println(*p)
	*p = 6
	fmt.Println(a)

	q := &a
	fmt.Println(q)
	fmt.Println(*q)
	*q = 7
	fmt.Println(a)
	fmt.Println(*p)
}

func Change1(a int) {
	a = 11
}

func Change2(p *int) {
	*p = 33
}

func Demo3() {
	a := 2

	Change1(a)
	fmt.Println(a)

	Change2(&a)
	fmt.Println(a)
}

//Xay dung ham Swap doi gia tri 2 tham so a, b kieu so nguyen, su dung
//con tro Swap(a,b)
func Swap(p, q *int) {
	tmp := *p
	*p = *q
	*q = tmp
}

func Demo4() {
	a := 1
	b := 5
	Swap(&a, &b)
	fmt.Println("a:", a, "b:", b)
}

func Modify1(a [3]int) {
	a[1] = 11
}

func Modify2(p *[3]int) {
	(*p)[1] = 22
}

func Demo() {
	var a = [3]int{1, 2, 3}
	Modify1(a)
	fmt.Println(a)

	Modify2(&a)
	fmt.Println(a)
}
