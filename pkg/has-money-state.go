package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("money already inserted")
}

func (h *HasMoneyState) DispenseItem() error {
	h.vendingMachine.itemCount--
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.SetState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.SetState(h.vendingMachine.hasItem)
	}
	return nil
}
