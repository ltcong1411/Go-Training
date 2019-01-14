package demo_json

import (
	"encoding/json"
	"fmt"
	"session7/entities"
)

func Demo1() {
	product := entities.Product{"p01", "name 1", 20, 30}
	str_json, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str_json))
	}
}

func Demo2() {
	str := `{"Id":"p01","Name":"name 1","Price":20,"Quantity":30}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)
	fmt.Println(product.ToString())
}

func Demo3() {
	str := `[
		{"Id":"p01","Name":"name 1","Price":20,"Quantity":30},
		{"Id":"p02","Name":"name 2","Price":30,"Quantity":20},
		{"Id":"p03","Name":"name 3","Price":40,"Quantity":50}
	]`
	var products []entities.Product
	json.Unmarshal([]byte(str), &products)
	for _, product := range products {
		fmt.Println(product.ToString())
		fmt.Println("---------------")
	}
}

func Demo4() {
	str := `{
		"Id": "st01", 
		"Name": "name 1", 
		"Address":{"Street": "street 1", "Ward": "ward 1"}
		}`

	var student entities.Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Student Info")
	fmt.Println("Id: ", student.Id)
	fmt.Println("Name: ", student.Name)
	fmt.Println("Street: ", student.Address.Street)
	fmt.Println("Ward: ", student.Address.Ward)

}

/*
[
	{
		"id":"cus1",
		"name":"invoice 1",
		"payment": "cash",
		"total":20,
		"customer": {
			"cus": "c01",
			"name": "name 1"
		},
		"status": "cancel",
		"products": [
			{

			}
		]
	},
	{
		"id":"cus2",
		"name":"invoice 2",
		"payment": "cash",
		"total":30,
		"customer": {
			"cus": "c02",
			"name": "name 2"
		},
		"status": "cancel"
	},
]
*/

func Demo() {
	str := `[
				{
					"id":"cus1",
					"name":"invoice 1",
					"payment": "cash",
					"total":20,
					"customer": {
						"id": "c01",
						"name": "name 1",
						"date_of_birthday": "1980-10-20"
					},
					"status": "cancel"
				},
				{
					"id":"cus2",
					"name":"invoice 2",
					"payment": "cash",
					"total":30,
					"customer": {
						"id": "c02",
						"name": "name 2",
						"date_of_birthday": "1985-10-20"
					},
					"status": "cancel"
				},
			]`
	var invoices []entities.Invoice
	err := json.Unmarshal([]byte(str), &invoices)
	if err != nil {
		fmt.Println(err)
	}

	for _, invoice := range invoices {
		fmt.Println(invoice.ToString())
		fmt.Println("---------------")
	}
}
