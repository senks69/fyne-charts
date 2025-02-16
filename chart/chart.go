package chart

import (
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Chart struct {
	widget.BaseWidget
	Points    []float64
	XLabels   []string
	YSteps    int
	LineColor color.Color
}

func NewChart(points []float64) *Chart {
	c := &Chart{
		Points:    points,
		YSteps:    5,
		LineColor: color.NRGBA{R: 0, G: 0, B: 255, A: 255},
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *Chart) AddValue(value float64) {
	c.Points = append(c.Points, value)
	c.YSteps = int(math.Ceil(value / 0.2))
	c.Refresh()
}

func (c *Chart) CreateRenderer() fyne.WidgetRenderer {
	return &chartRenderer{
		chart: c,
	}
}

type chartRenderer struct {
	chart   *Chart
	objects []fyne.CanvasObject
}

func (r *chartRenderer) Layout(size fyne.Size) {
	r.objects = []fyne.CanvasObject{}

	bg := canvas.NewRectangle(color.White)
	bg.Resize(size)
	r.objects = append(r.objects, bg)

	r.drawAxes(size)
	r.drawGrid(size)
	r.drawChartLine(size)
}

func (r *chartRenderer) drawAxes(size fyne.Size) {
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

func (r *chartRenderer) drawGrid(size fyne.Size) {
	gridColor := color.NRGBA{R: 200, G: 200, B: 200, A: 255}

	for i := 0; i < len(r.chart.Points); i++ {
		xPos := float32(i) * size.Width / float32(len(r.chart.Points)-1)

		line := canvas.NewLine(gridColor)
		line.Position1 = fyne.NewPos(xPos, 0)
		line.Position2 = fyne.NewPos(xPos, size.Height)
		r.objects = append(r.objects, line)

		if len(r.chart.XLabels) > i {
			label := canvas.NewText(r.chart.XLabels[i], color.Black)
			label.TextSize = 10
			label.Move(fyne.NewPos(xPos-5, size.Height+5))
			r.objects = append(r.objects, label)
		}
	}

	for i := 0; i <= r.chart.YSteps; i++ {
		yPos := size.Height * float32(i) / float32(r.chart.YSteps)

		line := canvas.NewLine(gridColor)
		line.Position1 = fyne.NewPos(0, yPos)
		line.Position2 = fyne.NewPos(size.Width, yPos)
		r.objects = append(r.objects, line)

		value := float64(i) / float64(r.chart.YSteps)
		text := canvas.NewText(strconv.FormatFloat(value, 'f', 1, 64), color.Black)
		text.TextSize = 10
		text.Move(fyne.NewPos(-25, yPos-5))
		r.objects = append(r.objects, text)
	}
}

func (r *chartRenderer) drawChartLine(size fyne.Size) {
	if len(r.chart.Points) < 2 {
		return
	}

	points := make([]fyne.Position, len(r.chart.Points))
	for i, val := range r.chart.Points {
		x := float32(i) * size.Width / float32(len(r.chart.Points)-1)
		y := size.Height - float32(val)*size.Height
		points[i] = fyne.NewPos(x, y)
	}

	for i := 1; i < len(points); i++ {
		line := canvas.NewLine(r.chart.LineColor)
		line.StrokeWidth = 2
		line.Position1 = points[i-1]
		line.Position2 = points[i]
		r.objects = append(r.objects, line)
	}
}

func (r *chartRenderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 300)
}

func (r *chartRenderer) Refresh() {
	r.Layout(r.chart.Size())
	canvas.Refresh(r.chart)
}

func (r *chartRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *chartRenderer) Destroy() {}
