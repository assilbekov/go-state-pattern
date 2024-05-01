package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) AddItem(count int) error {
	fmt.Println("Item already added, no need to add again")
	h.vendingMachine.itemCount += count
	return nil
}

func (h *HasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.SetState(h.vendingMachine.noItem)
		return fmt.Errorf("item out of stock")
	}
	h.vendingMachine.SetState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("item request pending")
}
