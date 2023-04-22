package main

import "fmt"

func main() {
	game := newGame()
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	dressFactorySingleInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactorySingleInstance.dressMap {
		fmt.Printf("dress color type: %s, dress color: %s\n", dressType, dress.getColor())
	}
}
