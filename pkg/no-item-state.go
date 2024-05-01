package pkg

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) AddItem(count int) error {
	n.vendingMachine.itemCount += count
	n.vendingMachine.SetState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}
