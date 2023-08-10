package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {

	app := app.New()
	calculator := NewCalculator()

	calculator.Build(app)

	app.Run()

}
