package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makeClockUI(app fyne.App) *widget.Button {
	return widget.NewButton("Clock", func() {
		openClockWindow(app)
	})
}

func openClockWindow(app fyne.App) {
	clockWindow := app.NewWindow("Clock")
	clockWindow.Resize(fyne.NewSize(200, 100))

	clock := widget.NewLabel("")
	updateTime(clock, time.Now())

	clockWindow.SetContent(clock)

	go func() {
		for t := range time.Tick(time.Second) {
			updateTime(clock, t)
		}
	}()

	clockWindow.Show()
}

func updateTime(clock *widget.Label, t time.Time) {
	s := t.Format("Time: 03:04:05")
	clock.SetText(s)
}
