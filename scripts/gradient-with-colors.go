package main

import (
	"image"
	"image/color"
	// "image/draw"
	"image/png"
	"math"
	"log"
	"os"
)

func drawGradientWithColors(img *image.RGBA, startColor, endColor color.RGBA, angle float64) {
	bounds := img.Bounds()
	w := bounds.Max.X
	h := bounds.Max.Y
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			var c color.RGBA
			if angle == 0 {
				c = interpolateColor(startColor, endColor, float64(x)/float64(w))
			} else {
				c = interpolateColor(startColor, endColor, (math.Sin(angle*float64(y))+1)/2)
			}
			img.Set(x, y, c)
		}
	}
}

func interpolateColor(start, end color.RGBA, t float64) color.RGBA {
	r := uint8(float64(start.R) + t*float64(end.R-start.R))
	g := uint8(float64(start.G) + t*float64(end.G-start.G))
	b := uint8(float64(start.B) + t*float64(end.B-start.B))
	a := uint8(float64(start.A) + t*float64(end.A-start.A))
	return color.RGBA{r, g, b, a}
}

func main() {
	width := 800
	height := 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	drawGradientWithColors(img, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, 0)
	
	f, err := os.Create("gradient.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
