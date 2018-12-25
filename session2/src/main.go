package main

import (
	// "mylib/lib1"
	// "mylib/lib2"
	// "mylib"

	"array/array1"
	"array/array2"
	f "fmt"
)

func main() {
	b := []int{5, 232, -325, -32, 23, 325}
	a := []int{0, -1, 1, 2, 3, 4, 5, 6, 11, 18, 7, 13}
	c := []int{0, -1, 1, 4}
	d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	e := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}
	f.Println(array1.SumAll(b))
	f.Println(array1.FindMaxMin(b))
	f.Println(array1.PrimeCount(a))
	f.Println(array1.ListPrime(a, 5, 13))
	f.Println(array1.HasPrime(c))
	f.Println(array2.SumAll(d))
	f.Println(array2.SumRow(d))
	f.Println(array2.HasPrime(d))
	f.Println(array2.HasPrime(e))
}

// func main() {
// 	// mylib.Hello()
// 	// f.Println(mylib.Hi("ABC"))

// 	// f.Println(lib1.show("chu dong tu"))
// 	// f.Println(lib2.display("nothing to display"))
// 	// f.Printf(" Area of a circle with radius = 3 is %f\n", circle.Area(3))
// 	// f.Printf(" Perimeter of a circle with radius = 3 is %f\n", circle.Per(3))
// 	// array.ArrayExported2()
// 	// a := []int{5, 232, 325, 32, 23, 325}
// 	// array.Array3(a)
// 	// array.Array4()
// 	// b := []int{5, 232, -325, -32, 23, 325}
// 	// f.Println(array.Array5(b))
// 	b := []int{5, 232, -325, -32, 23, 325}
// 	name := []string{"abbs", "asgsd", "asgd"}
// 	f.Println(array.Array6(b))
// 	array.Array7()
// 	array.Array8()
// 	array.Array9(b, name)
// 	f.Println(b)
// }
