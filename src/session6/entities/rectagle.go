package entities

type Rectagle struct {
	Name string
	A    float32
	B    float32
}

func (rectagle Rectagle) Area() float32 {
	return rectagle.A * rectagle.B
}

func (rectagle Rectagle) Perimeter() float32 {
	return (rectagle.A + rectagle.B) * 2
}

func (rectagle Rectagle) Type() string {
	return rectagle.Name
}
