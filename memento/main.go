package main

import "fmt"

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "a",
	}
	fmt.Printf("originator current state is %s\n", originator.state)
	caretaker.addMemento(originator.createMemento())

	originator.setState("b")
	fmt.Printf("originator current state is %s\n", originator.state)

	caretaker.addMemento(originator.createMemento())
	originator.setState("c")

	fmt.Printf("originator current state is %s\n", originator.state)
	caretaker.addMemento(originator.createMemento())

	originator.restorememento(caretaker.getMemento(1))
	fmt.Printf("restored to state %s\n", originator.getState())

	originator.restorememento(caretaker.getMemento(0))
	fmt.Printf("restored to state %s\n", originator.getState())
}
