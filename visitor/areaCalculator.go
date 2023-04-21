package main

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("calculating area for circle")
}

func (a *areaCalculator) visitForRectangle(s *rectangle) {
	fmt.Println("calculating area for rectangle")
}
