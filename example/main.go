package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senks69/fyne-charts/chart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chart Example")

	c := chart.NewBarChart([]float64{})
	c.XLabels = []string{}
	var startTime time.Time
	var endTime time.Duration
	isRunning := false
	var timeButton *widget.Button
	timeButton = widget.NewButton(
		"Start",
		func() {
			if isRunning {
				endTime = time.Since(startTime)
				c.AddValue(endTime.Seconds())
				fmt.Println(endTime.Seconds())
				timeButton.SetText("Start")
				isRunning = false
			} else {
				startTime = time.Now()
				timeButton.SetText("Stop")
				isRunning = true

			}
		},
	)
	w.SetContent(
		container.NewVBox(
			c,
			timeButton,
		),
	)

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
