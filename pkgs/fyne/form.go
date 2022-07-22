package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makeFormUI(app fyne.App) *widget.Button {
	return widget.NewButton("Form", func() {
		openFormWindow(app)
	})
}

func openFormWindow(app fyne.App) {
	formWindow := app.NewWindow("Form")
	formWindow.Resize(fyne.NewSize(500, 300))

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Name", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Name:", entry.Text)
			log.Println("Address:", textArea.Text)
			formWindow.Close()
		},
	}

	// we can also append items
	form.Append("Address", textArea)

	formWindow.SetContent(form)

	formWindow.Show()
}
