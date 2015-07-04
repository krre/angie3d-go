package core

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/krre/angie3d"
)

type Application struct {
	arguments []string
}

func NewApplication() *Application {
	application := Application{}

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	return &application
}

func (application *Application) Run() {
	defer glfw.Terminate()

	for !angie3d.IsExit {
		glfw.WaitEvents()
	}
}
