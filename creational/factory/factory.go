package main

import "fmt"

// IShape is the interface that defines methods common to all shapes.
type IShape interface {
	SetName(name string)
	GetName() string
	Draw() string
}

// Shape is a struct that implements the IShape interface.
type Shape struct {
	Name string
}

func (s *Shape) SetName(name string) {
	s.Name = name
}

func (s *Shape) GetName() string {
	return s.Name
}

func (s *Shape) Draw() string {
	return "Default draw method"
}

// Circle is a struct representing a circle, embedding the Shape struct.
type Circle struct {
	IShape
}

// NewCircle is a factory function to create a Circle instance.
func NewCircle() IShape {
	circle := &Circle{
		IShape: &Shape{},
	}
	circle.SetName("Circle")
	return circle
}

func (c *Circle) Draw() string {
	return "Drawing a circle"
}

// Square is a struct representing a square, embedding the Shape struct.
type Square struct {
	IShape
}

// NewSquare is a factory function to create a Square instance.
func NewSquare() IShape {
	square := &Square{
		IShape: &Shape{},
	}
	square.SetName("Square")
	return square
}

func (s *Square) Draw() string {
	return "Drawing a square"
}

// ShapeFactory is a simple factory that creates different types of shapes.
func ShapeFactory(shapeType string) (IShape, error) {
	switch shapeType {
	case "circle":
		return NewCircle(), nil
	case "square":
		return NewSquare(), nil
	default:
		return nil, fmt.Errorf("Unsupported shape type: %s", shapeType)
	}
}

func main() {
	circle, err := ShapeFactory("circle")
	if err != nil {
		fmt.Println(err)
	} else {
		printShapeDetails(circle)
	}

	square, err := ShapeFactory("square")
	if err != nil {
		fmt.Println(err)
	} else {
		printShapeDetails(square)
	}

	unknownShape, err := ShapeFactory("triangle")
	if err != nil {
		fmt.Println(err)
	} else {
		printShapeDetails(unknownShape)
	}
}

func printShapeDetails(shape IShape) {
	fmt.Printf("Shape Name: %s\n", shape.GetName())
	fmt.Println(shape.Draw())
	fmt.Println()
}
