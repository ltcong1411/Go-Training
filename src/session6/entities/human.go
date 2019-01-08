package entities

import (
	"fmt"
)

type IHuman1 interface {
	Eat() string
	Sleep()
}

type IHuman2 interface {
	Move()
}

type Human struct {
	Name, Gender string
}

func (human Human) ToString() string {
	return fmt.Sprintf("Name: %s\nGender: %s", human.Name, human.Gender)
}

func (human Human) Eat() string {
	return "human eat"
}

func (human Human) Sleep() {
	fmt.Println("human sleep")
}

func (human Human) Move() {
	fmt.Println("human move")
}
