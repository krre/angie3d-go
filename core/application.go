package core

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/krre/angie3d"
	"os"
)

type Application struct {
	arguments []string
}

func NewApplication() *Application {
	app := Application{}
	app.arguments = os.Args

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	return &app
}

func (app *Application) Run() {
	defer glfw.Terminate()

	for !angie3d.IsExit {
		glfw.WaitEvents()
	}
}

func (app *Application) GetArguments() []string {
	return app.arguments
}