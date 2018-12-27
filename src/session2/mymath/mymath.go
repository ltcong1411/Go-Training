package mymath

import "fmt"

func Add(a, b float32) float32 {
	return a + b
}

func Sub(a, b float32) float32 {
	return Add(a, -b)
}

func Mul(a, b float32) float32 {
	return a * b
}

func Divide(a, b float32) float32 {
	if b == 0.0 {
		fmt.Println("Denominator is zero!!! Zero is returned")
		return 0.0
	}
	return a / b
}
