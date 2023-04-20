package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) requestArrival() {
	if p.mediator.canLand(p) {
		fmt.Println("passengerTrain landing")
	} else {
		fmt.Println("passengerTrain waiting")
	}
}

func (p *passengerTrain) departure() {
	fmt.Println("passengerTrain leaving")
	p.mediator.notifyFree()
}
func (p *passengerTrain) permitArrival() {
	fmt.Println("passengerTrain arrival permitted, landing")
}
