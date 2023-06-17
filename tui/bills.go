package tui

import (
	"budget/dal"
	"github.com/rivo/tview"
	"strconv"
)

func AddBillAction() {
	newBill := dal.Bill{}
	billForm.Clear(true)

	billForm.AddInputField("Name", "", 20, nil, func(name string) {
		newBill.Name = name
	})
	billForm.AddInputField("Initial Balance", "0", 20, nil, func(balance string) {
		value, err := strconv.Atoi(balance)
		if err != nil {
			SimpleModal("Wrong value for Initial Balance", BillsList)
		}
		newBill.Balance = value
	})
	billForm.AddButton("Save", func() {
		err := dal.AddBill(newBill)
		if err != "" {
			SimpleModal(err, BillsList)
		}
		ClearFlex()
		getBills()
		pages.SwitchToPage(BillsList)
	})
}

func getBills() {
	bills := dal.GetBills()
	billsList.Clear()

	for index, bill := range bills {
		balance := strconv.Itoa(bill.Balance)
		billsList.AddItem(bill.Name+" Balance: "+balance, "", rune(49+index), nil)
	}

	billsListFlex.SetDirection(tview.FlexRow).
		AddItem(billsList, 0, 1, true).
		AddItem(billsListMenu, 0, 1, false)
}
