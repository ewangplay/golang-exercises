package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeGreetUI(app fyne.App) *widget.Button {
	return widget.NewButton("Greeting", func() {
		openGreetWindow(app)
	})
}

func openGreetWindow(app fyne.App) {
	greetWindow := app.NewWindow("Greeting")
	greetWindow.Resize(fyne.NewSize(200, 100))

	str := binding.NewString()

	greet := widget.NewLabelWithData(str)
	name := widget.NewEntry()
	name.SetPlaceHolder("Please input the name")

	name.OnChanged = func(content string) {
		str.Set(content)
	}

	greetWindow.SetContent(container.NewVBox(greet, name))

	greetWindow.Show()
}
