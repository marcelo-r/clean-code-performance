package enum

const (
	// Shape
	ShapeSquare = iota
	ShapeRectangle
	ShapeCircle
	ShapeTriangle
)

type Shape struct {
	shapeType int
	Height    float64
	Width     float64
}

func NewShape(shapeType int, height float64, width float64) *Shape {
	return &Shape{shapeType, height, width}
}

func (s *Shape) GetArea() float64 {
	switch s.shapeType {
	case ShapeSquare:
		return s.Height * s.Height
	case ShapeRectangle:
		return s.Height * s.Width
	case ShapeCircle:
		return 3.14 * s.Height * s.Height
	case ShapeTriangle:
		return s.Height * s.Width / 2
	default:
		return 0
	}
}

var Table = []float64{1, 1, 3.14, 0.5}

func (s *Shape) GetAreaTable() float64 {
	return Table[s.shapeType] * s.Width * s.Height
}
