package main

import "fmt"

type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("make ledger entry for accountID %s with txnType %s for amount %d\n", accountID, txnType, amount)
}
