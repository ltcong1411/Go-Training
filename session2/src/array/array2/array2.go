package array2

import "array/array1"

func SumAll(a [][]int) int {
	s := 0
	for r := 0; r < len(a); r++ {
		s += array1.SumAll(a[r])
	}
	return s
}

func SumRow(a [][]int) (b []int) {
	for r := 0; r < len(a); r++ {
		b = append(b, array1.SumAll(a[r]))
	}
	return
}

func HasPrime(a [][]int) bool {
	for r := 0; r < len(a); r++ {
		if array1.HasPrime(a[r]) {
			return true
		}
	}
	return false
}
