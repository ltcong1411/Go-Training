package entities

type Geomatry interface {
	Area() float32
	Perimeter() float32
	Type() string
}
