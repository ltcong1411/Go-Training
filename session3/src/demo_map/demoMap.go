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

func Demo4() {
	categories := map[string][]string{
		"categories1": []string{"prod1", "prod2", "prod3"},
		"categories2": []string{"prod4", "prod5", "prod6"},
		"categories3": []string{"prod7", "prod8", "prod9"},
	}
	categories["categories4"] = []string{"prod10", "prod11"}
	res := append(categories["categories2"], "prod6.1")
	fmt.Println(res)
	fmt.Println(categories)

	for key, value := range categories {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}
}
