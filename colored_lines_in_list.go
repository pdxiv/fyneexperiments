package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{
	"2021-09-22 09:23:12 INFO    Something happened, but it looks okay to me",
	"2021-09-22 09:53:32 WARNING Something mighty strange happpened",
	"2021-09-22 10:23:41 ERROR   Gosh dang! This is bad!",
	"2021-09-22 11:15:12 INFO    Nevermind, it's fine!",
}

var severityColor = []color.RGBA{
	{G: 127, A: 100},
	{R: 255, G: 127, A: 100},
	{R: 255, A: 100},
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.New(layout.NewMaxLayout(),
				canvas.NewRectangle(color.RGBA{G: 155, A: 100}),
				widget.NewLabel("Item x"))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			box := o.(*fyne.Container)
			rect := box.Objects[0].(*canvas.Rectangle)
			rect.FillColor = severityColor[i%len(severityColor)]
			label := box.Objects[1].(*widget.Label)
			label.TextStyle.Monospace = true
			label.SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
