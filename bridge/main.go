package main

import "fmt"

func main() {
	hpPrinter := &hp{}
	esponPrinter := epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()
	macComputer.setPrinter(&esponPrinter)
	macComputer.print()
	fmt.Println()
	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()
	winComputer.setPrinter(&esponPrinter)
	winComputer.print()
}
