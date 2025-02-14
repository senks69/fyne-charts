package main

import (
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senks69/fyne-charts/chart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Advanced Chart Example")

	c := chart.New([]float64{0.1, 0.4, 0.9, 0.2, 0.7})
	c.XLabels = []string{"Jan", "Feb", "Mar", "Apr", "May"}

	w.SetContent(
		container.NewVBox(
			c,
			widget.NewButton("Add Point", func() {
				c.Points = append(c.Points, rand.Float64())
				c.Refresh()
			}),
		),
	)

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
