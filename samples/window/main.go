package main

import (
	"github.com/krre/angie3d/core"
	"github.com/krre/angie3d/window"
)

func main() {
	app := core.NewApplication()

	window1 := window.NewWindow()
	window1.SetTitle("win 1")
	window1.Show()

	window2 := window.NewWindow()
	window2.SetTitle("win 2")
	window2.Show()

	app.Run()
}
