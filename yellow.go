package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var (
	background = color.CMYK{0, 0, 0x3f, 0x00}
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title: "Yellow Window",
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		var sz size.Event
		for {
			e := w.NextEvent()
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
				if e.Direction == key.DirPress {
					y := uint8(64 + rand.Intn(256-64))
					newyellow := color.CMYK{0, 0, y, 0}
					maxsize := sz.Size()
					rx := rand.Intn(maxsize.X)
					ry := rand.Intn(maxsize.Y)
					rsize := rand.Intn(maxsize.X) / 3
					w.Fill(image.Rect(rx, ry, rx+rsize, ry+rsize), newyellow, draw.Src)
					w.Publish()
				}
			case paint.Event:
				w.Fill(sz.Bounds(), yellow, draw.Src)
				w.Publish()
			case mouse.Event:
				y := uint8(64 + rand.Intn(256-64))
				newyellow := color.CMYK{0, 0, y, 0}
				maxsize := sz.Size()
				rx := rand.Intn(maxsize.X)
				ry := rand.Intn(maxsize.Y)
				rsize := rand.Intn(maxsize.X) / 3
				w.Fill(image.Rect(rx, ry, rx+rsize, ry+rsize), newyellow, draw.Src)
				w.Publish()
			case size.Event:
				sz = e
			case error:
				log.Print(e)
			}
		}
	})
}
