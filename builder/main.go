package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("normal house num floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("igloo house door type: %s\n", iglooHouse.doorType)
	fmt.Printf("igloo house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("igloo house num floor: %d\n", iglooHouse.floor)
}
