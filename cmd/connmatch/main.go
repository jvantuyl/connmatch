package main

import (
	"github.com/jvantuyl/connmatch/internal/app"
)

func main() {
	a := app.New()
	a.Init()
	a.Run()
}
