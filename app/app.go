package app

import "trading/config"

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Run() {
	// TODO start app logic
}

func (app *Application) Shutdown() {
	println("Off")
}
