package core

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"os"
	"runtime"
)

type Application struct {
	arguments []string
}

var (
	app    *Application
	isExit = false
)

func init() {
	runtime.LockOSThread()
}

func NewApplication() *Application {
	if app != nil {
		return app
	}

	app = new(Application)
	app.arguments = os.Args

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Visible, glfw.False)

	return app
}

func (app *Application) Run() {
	defer glfw.Terminate()

	for !isExit {
		glfw.WaitEvents()
	}
}

func (app *Application) GetArguments() []string {
	return app.arguments
}

func GetApplication() *Application {
	return app
}

func (app *Application) Exit() {
	isExit = true
}
