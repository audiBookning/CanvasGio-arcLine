package main

import (
	"image/color"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/op"

	"github.com/ajstarks/giocanvas"
)

const (
	// basic positions in radians
	right  = pi * (0.0 / 2.0)
	top    = pi * (1.0 / 2.0)
	left   = pi * (2.0 / 2.0)
	bottom = pi * (3.0 / 2.0)
	//
	pi          = math.Pi
	opFramerate = time.Second / 25.0 // 40ms
)

var (
	angle1                  = pi * 2.0
	angle3                  = pi * 2.0
	x1          float32     = 20.0
	y1          float32     = 80.0
	x3          float32     = 60.0
	y3          float32     = 80.0
	strokeColor color.NRGBA = color.NRGBA{128, 0, 0, 255}
)

type CanvasArc struct {
	*giocanvas.Canvas
}

func (c CanvasArc) ArcLine(x, y, radius float32, a1, a2 float64, size float32, fillcolor color.NRGBA) {
	// Ensure the angles are in the range [0, 2π)
	startAngle := math.Mod(a1, 2.0*pi)
	endAngle := math.Mod(a2, 2.0*pi)

	// Calculate lineSteps based on the radius
	// Smaller steps for larger radius
	lineSteps := float64(1.0 / (3.0 * radius * 2.0 * pi))

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
		endAngle += 2.0 * pi
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
	go func() {
		w := new(app.Window)
		if err := draw(w); err != nil {
			panic(err)
		}
	}()
	app.Main()
}

func draw(w *app.Window) error {

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// Calculate delta time
			currentStep := 0.1

			// canvas
			canvas := CanvasArc{giocanvas.NewCanvas(float32(e.Size.X), float32(e.Size.Y), app.FrameEvent{})}
			// context must be set before drawing to avoid bug?
			gtx := app.NewContext(canvas.Context.Ops, e)

			// Draw the arcs
			drawArcAntiClockwise(&canvas, currentStep)
			drawArcClockwise(&canvas, currentStep)

			// Redraw the canvas after opFramerate duration
			inv := op.InvalidateCmd{At: gtx.Now.Add(opFramerate)}
			gtx.Execute(inv)

			e.Frame(canvas.Context.Ops)
		}
	}
}

func drawArcAntiClockwise(canvas *CanvasArc, currentStep float64) {
	angle1 += currentStep
	// avoid angle1 getting too big for a float64 on long running animation
	angle1 = math.Mod(angle1, 2.0*pi)

	base := left
	canvas.ArcLine(x1, y1, 5, base, base+angle1, 0.2, strokeColor)
}

func drawArcClockwise(canvas *CanvasArc, currentStep float64) {
	angle3 += currentStep
	// avoid angle3 getting too big for a float64 on long running animation
	angle3 = math.Mod(angle3, 2.0*pi)

	base := top
	canvas.ArcLine(x3, y3, 5, base-angle3, base, 0.2, strokeColor)
}
