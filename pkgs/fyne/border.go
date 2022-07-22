package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeBorderUI(app fyne.App) *widget.Button {
	return widget.NewButton("Border Layout", func() {
		openBorderWindow(app)
	})
}

func openBorderWindow(app fyne.App) {
	borderWindow := app.NewWindow("Border Layout")
	borderWindow.Resize(fyne.NewSize(500, 500))

	top := container.NewCenter(canvas.NewText("top", color.Black))
	bottom := container.NewCenter(canvas.NewText("bottom", color.Black))
	left := canvas.NewText("left", color.Black)
	right := canvas.NewText("right", color.Black)
	middle := container.NewCenter(canvas.NewText("I am a happy tutu~", color.Black))

	content := container.NewBorder(top, bottom, left, right, middle)

	borderWindow.SetContent(content)

	borderWindow.Show()
}
