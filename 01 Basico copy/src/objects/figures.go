package objects

//cuadrilatero
type Cuadrilatero struct {
	Base   float64
	Altura float64
}

//circulo
type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Radio * c.Radio
}

func (c Cuadrilatero) Area() float64 {
	return c.Base * c.Altura
}
