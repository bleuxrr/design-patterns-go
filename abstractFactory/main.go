package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	printShoeDetail(adidasShoe)
	printShortDetail(adidasShort)
	printShoeDetail(nikeShoe)
	printShortDetail(nikeShort)
}

func printShoeDetail(s iShoe) {
	fmt.Printf("shoe logo: %s, size: %d\n", s.getLogo(), s.getSize())
}

func printShortDetail(s iShort) {
	fmt.Printf("short logo: %s, size: %d\n", s.getLogo(), s.getSize())
}
