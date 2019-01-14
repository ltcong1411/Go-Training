package entities

import "fmt"

type Student struct {
<<<<<<< HEAD
	Id, Name string
	Age      int
}

func (student Student) ToString() string {
	return fmt.Printf("id: %s\nname : %s\n age: %d", student.Id, student.Name, student.Age)
}

func NewStudent(id string, name string, age int) {
	student := Student(id, name, age)
	return student
}
=======
	Id   string
	Name string
	Age  int
}

func (student Student) ToString() string {
	return fmt.Sprintf("Id: %s, Name: %s, Age: %d", student.Id, student.Name, student.Age)
}

func NewStudent(id, name string, age int) Student {
	student := Student{id, name, age}
	return student
}

func (student Student) SetName(newName string) {
	student.Name = newName
}

func (student *Student) SetAge(newAge int) {
	student.Age = newAge
}
>>>>>>> b16df0e3461b618480359a36a8992fd0fd167abc
