package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("payment already done")
		return
	}
	fmt.Println("cashier getting money from patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}
