package training

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Circle, Area metodunu implement ediyor
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle, Area metodunu implement ediyor
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
