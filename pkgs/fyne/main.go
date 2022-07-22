package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	a := app.New()

	mainWindow := a.NewWindow("Main")
	mainWindow.Resize(fyne.NewSize(350, 600))
	mainWindow.SetMaster()

	// Set system tray
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				mainWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	mainWindow.SetContent(container.NewVBox(
		makeClockUI(a),
		makeGreetUI(a),
		makeGridWrapUI(a),
		makeBorderUI(a),
		makeFormUI(a),
		makeListUI(a)))

	// Just hide the main window when user selecting to close the main window
	mainWindow.SetCloseIntercept(func() {
		mainWindow.Hide()
	})

	mainWindow.Show()
	a.Run()
}
