package entities

type Student struct {
	Id      string `json:"id"` //trong json neu id thuong cung map dc
	Name    string `json:"name`
	Address Address
}

type Address struct {
	Street, Ward string
}
