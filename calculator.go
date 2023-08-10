package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	equation string

	input   *widget.Label
	buttons map[string]*widget.Button
	window  fyne.Window
}

func (calc *Calculator) updateText(text string) {
	calc.equation = calc.equation + text
	calc.input.SetText(calc.equation)
}

func (calc *Calculator) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	calc.buttons[text] = button

	return button
}

func (calc *Calculator) characterButton(character rune) *widget.Button {
	return calc.addButton(string(character), func() {
		calc.updateText(string(character))
	})

}

func (calc *Calculator) Build(app fyne.App) {
	calc.input = &widget.Label{Alignment: fyne.TextAlignCenter}
	calc.input.TextStyle.Monospace = true

	//Why does it give that error.

	calc.window = app.NewWindow("Calculator")
	calc.window.SetContent(container.NewGridWithColumns(1,
		calc.input,
		container.NewGridWithColumns(4,
			calc.characterButton('C'),
			calc.characterButton('('),
			calc.characterButton(')'),
			calc.characterButton('/')),
		container.NewGridWithColumns(4,
			calc.characterButton('7'),
			calc.characterButton('8'),
			calc.characterButton('9'),
			calc.characterButton('*')),
		container.NewGridWithColumns(4,
			calc.characterButton('4'),
			calc.characterButton('5'),
			calc.characterButton('6'),
			calc.characterButton('-')),
		container.NewGridWithColumns(4,
			calc.characterButton('1'),
			calc.characterButton('2'),
			calc.characterButton('3'),
			calc.characterButton('+')),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				calc.characterButton('0'),
				calc.characterButton('.')),
			calc.characterButton('='))),
	)

	//canvas := calc.window.Canvas()
	calc.window.Resize(fyne.NewSize(500, 500))
	calc.window.Show()
}

func NewCalculator() *Calculator {
	return &Calculator{
		buttons: make(map[string]*widget.Button, 19),
	}
}
