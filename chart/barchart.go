package chart

import (
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type BarChart struct {
	widget.BaseWidget
	Points   []float64
	XLabels  []string 
	YSteps   int      
	BarColor color.Color
}

func NewBarChart(points []float64) *BarChart {
	c := &BarChart{
		Points:   points,
		YSteps:   5,
		BarColor: color.NRGBA{R: 0, G: 0, B: 255, A: 255},
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *BarChart) AddValue(value float64){
	c.Points = append(c.Points, value)
	YSteps := int(math.Ceil(value / 0.2))
	if YSteps>c.YSteps{
		c.YSteps = int(value / 0.2)+2
	}
	c.Refresh()
}

func (c *BarChart) CreateRenderer() fyne.WidgetRenderer {
	return &BarChartRenderer{
		chart: c,
	}
}

type BarChartRenderer struct {
	chart   *BarChart
	objects []fyne.CanvasObject
}

func (r *BarChartRenderer) Layout(size fyne.Size) {
	r.objects = []fyne.CanvasObject{}

	bg := canvas.NewRectangle(color.White)
	bg.Resize(size)
	r.objects = append(r.objects, bg)

	r.drawAxes(size)
	r.drawGrid(size)
	r.drawBars(size)
}

func (r *BarChartRenderer) drawAxes(size fyne.Size) {
	axisColor := color.Black

	xAxis := canvas.NewLine(axisColor)
	xAxis.Position1 = fyne.NewPos(0, size.Height)
	xAxis.Position2 = fyne.NewPos(size.Width, size.Height)
	r.objects = append(r.objects, xAxis)

	yAxis := canvas.NewLine(axisColor)
	yAxis.Position1 = fyne.NewPos(0, 0)
	yAxis.Position2 = fyne.NewPos(0, size.Height)
	r.objects = append(r.objects, yAxis)

	xLabel := canvas.NewText("Ось X", color.Black)
	xLabel.TextSize = 12
	xLabel.Move(fyne.NewPos(size.Width-40, size.Height-20))

	yLabel := canvas.NewText("Ось Y", color.Black)
	yLabel.TextSize = 12
	yLabel.Move(fyne.NewPos(10, 5))

	r.objects = append(r.objects, xLabel, yLabel)
}

func (r *BarChartRenderer) drawGrid(size fyne.Size) {
	gridColor := color.NRGBA{R: 200, G: 200, B: 200, A: 255}

	maxValue := 0.0
	for _, val := range r.chart.Points {
		if val > maxValue {
			maxValue = val
		}
	}

	for i := 0; i <= r.chart.YSteps; i++ {
		yPos := size.Height * float32(i) / float32(r.chart.YSteps)

		line := canvas.NewLine(gridColor)
		line.Position1 = fyne.NewPos(0, yPos)
		line.Position2 = fyne.NewPos(size.Width, yPos)
		r.objects = append(r.objects, line)

		value := maxValue * float64(i) / float64(r.chart.YSteps)
		text := canvas.NewText(strconv.FormatFloat(value, 'f', 1, 64), color.Black)
		text.TextSize = 10
		text.Move(fyne.NewPos(-25, yPos-5))
		r.objects = append(r.objects, text)
	}
}

func (r *BarChartRenderer) drawBars(size fyne.Size) {
	if len(r.chart.Points) == 0 {
		return
	}

	barWidth := size.Width / float32(len(r.chart.Points)) * 0.8
	spacing := size.Width / float32(len(r.chart.Points)) * 0.1

	for i, val := range r.chart.Points {
		x := float32(i)*(barWidth+spacing) + spacing
		height := float32(val) * size.Height

		bar := canvas.NewRectangle(r.chart.BarColor)
		bar.Resize(fyne.NewSize(barWidth, height))
		bar.Move(fyne.NewPos(x, size.Height-height))

		label := canvas.NewText(strconv.FormatFloat(val, 'f', 2, 64), color.Black)
		label.TextSize = 10
		label.Move(fyne.NewPos(x+barWidth/2-10, size.Height-height-15))

		r.objects = append(r.objects, bar, label)
	}
}

func (r *BarChartRenderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 300)
}

func (r *BarChartRenderer) Refresh() {
	r.Layout(r.chart.Size())
	canvas.Refresh(r.chart)
}

func (r *BarChartRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *BarChartRenderer) Destroy() {}