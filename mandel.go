package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/cmplx"
	"os"
)

func main() {
	var renderer *sdl.Renderer

	var width int = 320
	var height int = 240

	hwidth := int(width / 2)
	hheight := int(height / 2)

	var wid int = 4
	var xcenter, ycenter int = -1, 0

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Mandelbrot", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
	defer renderer.Destroy()

	renderer.Clear()
	renderer.Present()

	var c1, c2, c3 uint8 = 0, 0, 0

	for xcoord := 0; xcoord < width; xcoord++ {
		for ycoord := 0; ycoord < height-1; ycoord++ {

			fmt.Printf("%d %d\n", xcoord, ycoord)

			ca := (xcoord-hwidth)/width*wid + xcenter
			cb := (ycoord-hheight)/width*1*wid + ycenter

			res, i := mandelbrot(complex(float64(ca), float64(cb)))

			var hcolor uint8 = 128

			if i != 0 {
				hcolor = uint8(10 * i)
			}

			if res == 0 {
				c1 = 0
				c2 = 0
				c3 = 0
			} else if i < 5 {
				c1 = 0
				c2 = 0
				c3 = 128
			} else if i > 5 && i < 7 {
				c1 = 0
				c2 = 0
				c3 = hcolor
			} else if i > 7 && i < 10 {
				c1 = 0
				c2 = hcolor
				c3 = 0
			} else {
				c1 = hcolor
				c2 = 0
				c3 = 0
			}

			renderer.SetDrawColor(c1, c2, c3, 0)
			renderer.DrawPoint(xcoord, ycoord)
			renderer.Present()
		}

	}

	sdl.Delay(5000)
	sdl.Quit()
}

func mandelbrot(c complex128) (complex128, int) {
	var z complex128 = c
	var i int
	for i = 1; i < 21; i++ {
		z = z*z + c
	}

	if cmplx.Abs(z) > 2 {
		return z, i
	}
	return 0, 0
}
