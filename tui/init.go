package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var App = tview.NewApplication()
var pages = tview.NewPages()
var modal = tview.NewModal().SetText("")

var billForm = tview.NewForm()
var billsList = tview.NewList()
var billsListFlex = tview.NewFlex()

func SimpleModal(text string, switchTo string) {
	modal.SetText(text)
	pages.SwitchToPage(Modal)
	modal.ClearButtons()
	modal.AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				pages.SwitchToPage(switchTo)
			}
		})
}

func ClearFlex() {
	billsListFlex.Clear()
}

func InitTui() {
	pages.AddPage(MainMenu, mainMenu, true, true)
	pages.AddPage(BillsList, billsListFlex, true, false)
	pages.AddPage(AddBill, billForm, true, false)
	pages.AddPage(Modal, modal, true, false)

	App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 113: // q
			App.Stop()
		case 109: // m
			ClearFlex()
			pages.SwitchToPage(MainMenu)
		case 97: // a
			ClearFlex()
			pageName, _ := pages.GetFrontPage()
			if pageName == BillsList {
				pages.SwitchToPage(AddBill)
				AddBillAction()
			}
		case 98: // b
			ClearFlex()
			getBills()
			pages.SwitchToPage(BillsList)
		}
		return event
	})

	if err := App.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
