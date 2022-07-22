package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeGridWrapUI(app fyne.App) *widget.Button {
	return widget.NewButton("Grid Wrap Layout", func() {
		openGridWindow(app)
	})
}

func openGridWindow(app fyne.App) {
	gridWindow := app.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("Tom", color.Black)
	text2 := canvas.NewText("Mary", color.Black)
	text3 := canvas.NewText("Samy", color.Black)
	text4 := canvas.NewText("Luffy", color.Black)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3, text4)

	gridWindow.SetContent(grid)

	gridWindow.Show()
}
