package main

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restorememento(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}
