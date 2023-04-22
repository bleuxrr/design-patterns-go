package main

import "fmt"

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("sending wallet debit notification")
}
