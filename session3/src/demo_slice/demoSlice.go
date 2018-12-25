package demo_slice

import (
	"fmt"
	"sort"
)

func Demo1() {
	var a = [5]int{5, 1, -1, 9, 12}
	fmt.Println(a)
	a1 := a[0:3]
	fmt.Println(a1)
	a1[1] = 999
	fmt.Println(a)
	a2 := a[0:2]
	a2[1] = 777
	fmt.Println(a)
	fmt.Println(a2)
}

func Demo2() {
	var a = [5]int{5, 1, -1, 9, 12}
	fmt.Println(a)
	a1 := a[:2]
	fmt.Println(a1)
	a2 := a[2:]
	fmt.Println(a2)
	a3 := a[:]
	fmt.Println(a3)
}

func Demo3() {
	var a = [5]int{5, 1, -1, 9, 12}
	a1 := a[1:4]
	fmt.Println(a1)
	fmt.Println("len: ", len(a1))
	fmt.Println("cap: ", cap(a1))

	a1 = append(a1, 20)
	fmt.Println("append 20")
	fmt.Println("len: ", len(a1))
	fmt.Println("cap: ", cap(a1))

	a1 = append(a1, 30)
	fmt.Println("append 30")
	fmt.Println("len: ", len(a1))
	fmt.Println("cap: ", cap(a1))
}

func Demo4() {
	a := make([]int, 3, 4)
	a[0] = 1
	a[1] = 5
	a[2] = 10
	fmt.Println(a)
}

func Change1(a [5]int) {
	a[1] = 999
}

func Change2(a []int) {
	a[1] = 888
}

func Demo5() {
	var a = [5]int{5, 1, -1, 9, 12}
	fmt.Println(a)
	Change1(a)
	fmt.Println(a)

	Change2(a[:])
	fmt.Println(a)
}

func Demo6() {
	var a = [5]int{5, 1, -1, 9, 12}
	b := a[:]
	sort.Slice(b, func(i, j int) bool {
		return b[i] <= b[j]
	})
	fmt.Println(a)
}
