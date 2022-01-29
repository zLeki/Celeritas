package main

import (
	"github.com/zLeki/Celeritas"
	"os"
)

func initApplication() *App {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		return nil
	}
	cel.AppName = "MyApp"
	cel.Infolog.Println("Debug is set to,", cel.Debug)
	app := &App{
		App: cel,
	}
	return app
}
