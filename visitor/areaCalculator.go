package main

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("calculating area for square ", s.side*s.side)
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("calculating area for circle ", 3.14*float64(c.radius*c.radius))
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	fmt.Println("calculating area for rectangle ", r.l*r.b)
}
