package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("goodsTrain landing")
	} else {
		fmt.Println("goodsTrain waiting")
	}
}

func (g *goodsTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("goodsTrain leaving")
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("goodsTrain arrival permitted, landing")
}
