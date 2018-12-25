package demo_map

import (
	"fmt"
)

func Demo1() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name1",
		"address": "address 1",
	}
	student["phone"] = "1234"
	fmt.Println(student)
	for key, value := range student {
		fmt.Println(key, value)
	}
}

func Demo2() {
	student := make(map[string]string)
	student["id"] = "st01"
	student["name"] = "name1"
	student["email"] = "abc@gmail.com"
	fmt.Println(student)
}

func Demo3() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name1",
		"address": "address 1",
	}
	value, ok := student["name"]
	fmt.Println(value)
	fmt.Println(ok)

	delete(student, "address")
	fmt.Println(student)
}
