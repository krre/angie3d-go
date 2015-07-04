package window

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/krre/angie3d/core"
	"image/color"
)

type Window struct {
	color   color.Color
	handler *glfw.Window
	height  int
	title   string
	width   int
	x       int
	y       int
}

var (
	winList  []*Window
	winCount = 0
)

func NewWindow() *Window {
	window := Window{
		color:  color.Gray{},
		height: 480,
		title:  "Angie3D",
		width:  640}

	winList = append(winList, &window)
	winCount++

	handler, err := glfw.CreateWindow(window.width, window.height, window.title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.handler = handler

	handler.SetCloseCallback(func(w *glfw.Window) {
		w.Destroy()
		winCount--
		if winCount == 0 {
			core.GetApplication().Exit()
		}
	})

	return &window
}

func (window *Window) SetTitle(title string) {
	window.title = title
	window.handler.SetTitle(title)
}

func (window *Window) Show() {
	window.handler.Show()
}

func (window *Window) Hide() {
	window.handler.Hide()
}
