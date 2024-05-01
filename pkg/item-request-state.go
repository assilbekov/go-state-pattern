package pkg

import "fmt"

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestState) AddItem(count int) error {
	return fmt.Errorf("item request in progress")
}

func (i *ItemRequestState) RequestItem() error {
	return fmt.Errorf("item request in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less")
	}
	i.vendingMachine.SetState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("payment pending")
}
