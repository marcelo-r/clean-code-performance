package polymorphism

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func NewSquare(l float64) Shape {
	return &Square{l}
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(w, h float64) Shape {
	return &Rectangle{w, h}
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) Shape {
	return &Circle{r}
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

type Triangle struct {
	Base   float64
	Height float64
}

func NewTriangle(b, h float64) Shape {
	return &Triangle{b, h}
}

func (t *Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
