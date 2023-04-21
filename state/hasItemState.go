package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item stored")
	}
	fmt.Printf("item requested\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequsted)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}
