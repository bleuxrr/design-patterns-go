package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) addItem(_ int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("dispensing item")
	i.vendingMachine.itemCount--
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState((i.vendingMachine.hasItem))
	}
	return nil
}
