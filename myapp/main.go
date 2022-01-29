package main

import (
	"github.com/zLeki/Celeritas"
)

type App struct {
	App *celeritas.Celeritas
}

func main() {
	c := initApplication()
	err := c.App.ListenAndServe()
	if err != nil {
		return
	}
}
