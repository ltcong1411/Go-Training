package entities

import "fmt"

/*
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
*/

type Invoice struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Payment  string   `json:"payment"`
	Total    int      `json:"total"`
	Customer Customer `json:"customer"`
	Status   string   `json:"status"`
}

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Birth string `json:"date_of_birthday"`
}

func (invoice Invoice) ToString() string {
	return fmt.Sprintf("Id: %s\nName: %s\nPayment: %s\nTotal: %d\nCustomerId: %s\nCustomerName: %s\nBirthday: %s", invoice.Id, invoice.Name, invoice.Payment, invoice.Total, invoice.Customer.Id, invoice.Customer.Name, invoice.Customer.Birth)
}
