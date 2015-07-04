package main

import (
	"github.com/krre/angie3d/core"
	"github.com/krre/angie3d/window"
)

func main() {
	app := core.NewApplication()

	win := window.NewWindow()
	win.SetTitle("Hello World")
	win.Show()

	app.Run()
}
