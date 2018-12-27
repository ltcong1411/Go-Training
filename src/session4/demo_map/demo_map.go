package demo_map

import "fmt"

func Demo1() {
	student := make(map[string]string)
	student["id"] = "st01"
	student["name"] = "name 1"
	student["address"] = "address 1"
	fmt.Println(student)
	fmt.Println("Student Info")
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func Demo2() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	student["phone"] = "123456"
	fmt.Println(student)
}

func Demo3() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	value, ok := student["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}

	delete(student, "address")
	fmt.Println(student)
}

func Display(student map[string]string) {
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func Demo4() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}

	Display(student)
}

func Demo() {
	categories := map[string][]string{
		"category1": []string{"prod 1", "prod 2"},
		"category2": []string{"prod 3", "prod 4", "prod 5"},
		"category3": []string{"prod 6", "prod 7", "prod 8"},
	}

	fmt.Println(categories)
	categories["category4"] = []string{"prod 9", "prod 10", "prod 11"}
	fmt.Println(categories)
	categories["category3"] = append(categories["category3"], "prod 5.1", "prod 5.2")
	fmt.Println(categories)

	for key, value := range categories {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println(v)
		}
	}
}
