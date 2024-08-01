package main

import (
	"image/color"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"

	"github.com/ajstarks/giocanvas"
)

const (
	// basic positions in radians
	right  = 0
	top    = math.Pi / 2
	left   = math.Pi
	bottom = math.Pi * 1.5
	twoPi  = math.Pi * 2

	// arc coordinates
	x1  = 25.0
	mid = 50.0
	x3  = 100 - x1

	opFramerate = time.Second / 50.0 // 40ms
)

type CanvasArc struct {
	*giocanvas.Canvas
}

func newArcLine(c *giocanvas.Canvas, x, y, radius float32, a1, a2 float64, size float32, fillcolor color.NRGBA) {
	// Ensure the angles are in the range [0, 2π)
	startAngle := math.Mod(a1, twoPi)
	endAngle := math.Mod(a2, twoPi)

	// Calculate lineSteps based on the radius
	// Smaller steps for larger radius
	lineSteps := float64(1.0 / (3.0 * radius * twoPi))

	// Define minimum and maximum step sizes
	const minStepSize = 0.001 // Minimum allowed step size
	const maxStepSize = 0.1   // Maximum allowed step size

	// Clamp lineSteps to be within the defined range for performance reasons
	if lineSteps < minStepSize {
		lineSteps = minStepSize
	}
	if lineSteps > maxStepSize {
		lineSteps = maxStepSize
	}

	// Ensure we handle crossing the 0/2π boundary correctly
	if endAngle < startAngle {
		endAngle += twoPi
	}

	// Initialize the starting point
	x1, y1 := c.Polar(x, y, radius, float32(startAngle))

	for t := startAngle; t < endAngle; t += lineSteps {
		x2, y2 := c.Polar(x, y, radius, float32(t))
		c.Line(x1, y1, x2, y2, size, fillcolor)
		x1 = x2
		y1 = y2
	}
}

func main() {
	go draw()
	app.Main()
}

func draw() {
	w := new(app.Window)
	c1 := color.NRGBA{128, 0, 0, 255}
	c2 := color.NRGBA{0, 0, 128, 255}
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			// Calculate delta time
			currentStep := 0.1

			// canvas
			canvas := giocanvas.NewCanvas(float32(e.Size.X), float32(e.Size.Y), app.FrameEvent{})
			// context must be set before drawing to avoid bug?
			gtx := app.NewContext(canvas.Context.Ops, e)

			// Draw the arcs
			drawArcAntiClockwise(canvas, currentStep, c1)
			drawArcClockwise(canvas, currentStep, c2)

			// Redraw the canvas after opFramerate duration
			inv := op.InvalidateCmd{At: gtx.Now.Add(opFramerate)}
			gtx.Execute(inv)

			e.Frame(canvas.Context.Ops)
		}
	}
}

var (
	angle1 = twoPi
	angle3 = twoPi
)

func drawArcAntiClockwise(canvas *giocanvas.Canvas, currentStep float64, color color.NRGBA) {
	angle1 += currentStep
	// avoid angle1 getting too big for a float64 on long running animation
	angle1 = math.Mod(angle1, twoPi)

	base := top
	a1 := base
	a2 := base + angle1
	newArcLine(canvas, x1, mid, 5, a1, a2, 0.5, color)
	newArcLine(canvas, x1, mid, 10, a1, a2, 0.5, color)
	newArcLine(canvas, x1, mid, 20, a1, a2, 0.5, color)
}

func drawArcClockwise(canvas *giocanvas.Canvas, currentStep float64, color color.NRGBA) {
	angle3 += currentStep
	// avoid angle3 getting too big for a float64 on long running animation
	angle3 = math.Mod(angle3, twoPi)

	base := top
	a1 := base - angle3
	a2 := base
	newArcLine(canvas, x3, mid, 5, a1, a2, 0.5, color)
	newArcLine(canvas, x3, mid, 10, a1, a2, 0.5, color)
	newArcLine(canvas, x3, mid, 20, a1, a2, 0.5, color)
}

func stdDrawArcAntiClockwise(canvas *giocanvas.Canvas, currentStep float64, color color.NRGBA) {
	angle1 += currentStep
	// avoid angle1 getting too big for a float64 on long running animation
	angle1 = math.Mod(angle1, twoPi)

	base := top
	a1 := base
	a2 := base + angle1
	canvas.ArcLine(x1, mid, 5, a1, a2, 0.5, color)
	canvas.ArcLine(x1, mid, 10, a1, a2, 0.5, color)
	canvas.ArcLine(x1, mid, 20, a1, a2, 0.5, color)
}

func stdDrawArcClockwise(canvas *giocanvas.Canvas, currentStep float64, color color.NRGBA) {
	angle3 += currentStep
	// avoid angle3 getting too big for a float64 on long running animation
	angle3 = math.Mod(angle3, twoPi)

	base := top
	a1 := base - angle3
	a2 := base
	canvas.ArcLine(x3, mid, 5, a1, a2, 0.5, color)
	canvas.ArcLine(x3, mid, 10, a1, a2, 0.5, color)
	canvas.ArcLine(x3, mid, 20, a1, a2, 0.5, color)
}
