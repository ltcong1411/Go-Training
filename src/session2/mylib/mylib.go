package mylib

import "fmt"

func Hello() {
	fmt.Println("hello")
}

func Hi(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
