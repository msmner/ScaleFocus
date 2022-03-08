package main

import (
	"fmt"
	"math"
)

func main() {
	circleShape := Circle{}
	circleOne := circleShape.NewCircle(2)
	circleTwo := circleShape.NewCircle(1.2)

	squareShape := Square{}
	squareOne := squareShape.NewSquare(1.5)
	squareTwo := squareShape.NewSquare(4)

	shapes := Shapes{}
	shapes = append(shapes, circleOne, circleTwo, squareOne, squareTwo)

	maxArea := shapes.LargestArea()

	fmt.Printf("Larges area is: %f", maxArea)
}

type Shapes []Shape

type Shape interface {
	Area() float64
}

func (s Shapes) LargestArea() float64 {
	maxArea := s[0].Area()
	for _, shape := range s {
		if maxArea < shape.Area() {
			maxArea = shape.Area()
		}
	}

	return maxArea
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (circle *Circle) NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (square *Square) NewSquare(side float64) *Square {
	return &Square{side}
}

func (circle *Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (square *Square) Area() float64 {
	return math.Pow(square.side, 2)
}
