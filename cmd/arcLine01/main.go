// test arcs
package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"math"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/ajstarks/giocanvas"
)

const pi = math.Pi

func main() {
	var cw, ch int
	flag.IntVar(&cw, "width", 600, "canvas width")
	flag.IntVar(&ch, "height", 800, "canvas height")
	flag.Parse()
	width := float32(cw)
	height := float32(ch)

	go func() {
		w := &app.Window{}
		w.Option(app.Title("Arc"), app.Size(unit.Dp(width), unit.Dp(height)))
		if err := arc(w); err != nil {
			io.WriteString(os.Stderr, "Cannot create the window\n")
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

func arc(w *app.Window) error {

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			canvas := giocanvas.NewCanvas(float32(e.Size.X), float32(e.Size.Y), e)

			var x float32 = 50
			var y float32 = 50
			var radius float32 = 5
			var size float32 = 0.2
			var strokeColor color.NRGBA = color.NRGBA{128, 0, 0, 255}

			angle1 := pi

			// Ref: https://stackoverflow.com/a/22185792/3658510
			epsilon := math.Nextafter(1, 2) - 1
			fmt.Println("epsilon: ", epsilon)

			angle2 := angle1 + epsilon

			canvas.ArcLine(x, y, radius, angle1, angle2, size, strokeColor)

			e.Frame(canvas.Context.Ops)
		}
	}
}
