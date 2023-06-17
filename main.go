package main

import (
	"budget/dal"
	"budget/tui"
)

func main() {
	dal.ApplyMigrations()
	tui.InitTui()
}
