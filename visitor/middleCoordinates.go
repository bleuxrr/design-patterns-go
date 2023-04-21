package main

import "fmt"

type middleCoordinates struct {
	area int
}

func (m *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("calculating middle point coordinates for square")
}

func (m *middleCoordinates) visitForCircle(s *circle) {
	fmt.Println("calculating middle point coordinates for circle")
}

func (m *middleCoordinates) visitForRectangle(s *rectangle) {
	fmt.Println("calculating middle point coordinates for rectangle")
}
