package entities

import "fmt"

type Employee struct {
	Human  Human
	Id     string
	Salary float32
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("%s\nId: %s\nSalary: %f", employee.Human.ToString(), employee.Id, employee.Salary)
}
