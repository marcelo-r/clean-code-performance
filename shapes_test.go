package main

import (
	"testing"

	"github.com/marcelo-r/clean-code-performance/enum"
	"github.com/marcelo-r/clean-code-performance/polymorphism"
)

func areaPolymorphismList(n int) []polymorphism.Shape {
	shapes := make([]polymorphism.Shape, 0)
	for i := 0; i < n; i++ {
		f := float64(i)
		shapes = append(shapes, polymorphism.NewSquare(f))
		shapes = append(shapes, polymorphism.NewRectangle(f, f*2))
		shapes = append(shapes, polymorphism.NewCircle(f))
		shapes = append(shapes, polymorphism.NewTriangle(f, f))
	}
	return shapes
}

func areaEnumList(n int) []*enum.Shape {
	shapes := make([]*enum.Shape, 0)
	for i := 0; i < n; i++ {
		f := float64(i)
		shapes = append(shapes, enum.NewShape(enum.ShapeSquare, f, f))
		shapes = append(shapes, enum.NewShape(enum.ShapeRectangle, f, f*2))
		shapes = append(shapes, enum.NewShape(enum.ShapeCircle, f, 1))
		shapes = append(shapes, enum.NewShape(enum.ShapeTriangle, f, f))
	}
	return shapes
}

var (
	shapesNewInterface1    = areaPolymorphismList(1)
	shapesEnum1            = areaEnumList(1)
	shapesNewInterface1000 = areaPolymorphismList(1000)
	shapesEnum1000         = areaEnumList(1000)
)

func areaInterface(shapes []polymorphism.Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func areaEnum(shapes []*enum.Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.GetArea()
	}
	return area
}

func areaTable(shapes []*enum.Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.GetAreaTable()
	}
	return area
}

func BenchmarkAreaPoly1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaInterface(shapesNewInterface1)
	}
}

func BenchmarkAreaEnum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaEnum(shapesEnum1)
	}
}

func BenchmarkAreaTable1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaTable(shapesEnum1)
	}
}

func BenchmarkAreaPoly1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaInterface(shapesNewInterface1000)
	}
}

func BenchmarkAreaEnum1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaEnum(shapesEnum1000)
	}
}

func BenchmarkAreaTable1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaTable(shapesEnum1000)
	}
}
