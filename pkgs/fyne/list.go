package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeListUI(app fyne.App) *widget.Button {
	return widget.NewButton("List", func() {
		openListWindow(app)
	})
}

func openListWindow(app fyne.App) {
	listWindow := app.NewWindow("List")
	listWindow.Resize(fyne.NewSize(500, 500))

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("hello")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	list.OnSelected = func(id widget.ListItemID) {
		v, err := data.GetValue(id)
		if err != nil {
			return
		}
		fmt.Println(v)
	}

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	listWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))

	listWindow.Show()
}
