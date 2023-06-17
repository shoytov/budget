package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var mainMenu = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(m) Main Menu \n(b) switch to bills list \n(q) to quit")

var billsListMenu = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(m) Main Menu \n(a) Add a new bill\n(q) to quit")
