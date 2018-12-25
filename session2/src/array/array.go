package array

import (
	"fmt"
	f "fmt"
	"sort"
)

func ArrayExported1() {
	var a [5]int
	a[0] = 0
	a[3] = 1
	a[4] = 2
	a[2] = 3
	a[1] = 4
	f.Println(a)
	f.Println("size: ", len(a))

	for i := 0; i < len(a); i++ {
		f.Print("\t", a[i])
	}
}

func ArrayExported2() {
	a := []int{5, 232, 325, 32, 23, 325}
	f.Println(a)
	f.Println("size: ", len(a))
}

func Array3(a []int) {
	f.Println(a)
}

func Array4() {
	name := []string{"abbs", "asgsd", "asgd"}
	f.Println(name)

	name = append(name, "nodat")

	f.Println(name)
}

func Array5(a []int) []int {
	b := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			b = append(b, a[i])
		}
	}
	return b
}

func Array6(a []int) (b, c, d, e []int) {
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			b = append(b, a[i])
		} else {
			c = append(c, a[i])
		}

		if a[i]%2 == 0 {
			d = append(d, a[i])
		} else {
			e = append(e, a[i])
		}
	}
	return
}

func Array7() {
	name := []string{"abbs", "asgsd", "asgd"}
	for _, a := range name {
		f.Println(a)
	}
}

func Array8() {
	a := [2][3]int{
		{5, 3, 3},
		{3, 3, 3},
	}
	f.Println(a)
	f.Println("size : ", len(a))

	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a[r]); c++ {
			f.Print("\t", a[r][c])
		}
		f.Println()
	}
}

func Array9(a []int, b []string) {
	sort.Ints(a)
	fmt.Println(a)
	sort.Strings(b)
	fmt.Println(b)
}
