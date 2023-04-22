package main

import "fmt"

type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("printing by a hp printer")
}
