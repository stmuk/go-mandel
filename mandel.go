package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	start := time.Now()
	var renderer *sdl.Renderer

	var width int32 = 1920
	var height int32 = 1080

	hwidth := int32(width / 2)
	hheight := int32(height / 2)

	var wid float64 = 4
	var xcenter, ycenter float64 = -1, 0

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

	renderer.SetDrawColor(0, 0, 128, 0)

	renderer.Clear()
	renderer.Present()

	var c1, c2, c3 uint8 = 0, 0, 0

	var xcoord int32
	var ycoord int32

	for xcoord = 0; xcoord < width; xcoord++ {
		for ycoord = 0; ycoord < height-1; ycoord++ {

			ca := float64(xcoord-hwidth)/float64(width)*wid + xcenter
			cb := float64(ycoord-hheight)/float64(width)*1*wid + ycenter

			res, i := mandelbrot(complex(ca, cb))

			var hcolor uint8 = 128

			if i != 0 {
				hcolor = uint8(10 * i)
			}

			if res == 0 {
				c1 = 0
				c2 = 0
				c3 = 128
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
		}

	}

	renderer.Present()
	elapsed := time.Since(start)

	log.Printf("took %s sec(s)", elapsed)

	sdl.Delay(5000)
	sdl.Quit()
}

func mandelbrot(c complex128) (complex128, int) {
	var z = c
	var i int
	for i = 1; i < 20; i++ {
		z = z*z + c

		if cmplx.Abs(z) > 2 {
			return z, i
		}
	}
	return 0, 0
}
