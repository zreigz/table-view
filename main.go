package main

import (
	"github.com/rivo/tview"
	"runtime/debug"
)

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		debug.PrintStack()
		panic(err)
	}
}
