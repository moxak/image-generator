package main

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    "os"
    "log"
    "math"

    "golang.org/x/image/draw"
)


func drawGradient(width, height int, angle float64, startColor, midColor, endColor color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{startColor}, image.ZP, draw.Src)

	dx := float64(width) * math.Cos(angle*math.Pi/180)
	dy := float64(height) * math.Sin(angle*math.Pi/180)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// calculate linear interpolation between startColor, midColor and endColor
			// based on position in the gradient
			gradientPos := ((float64(x)*dx + float64(y)*dy) / (dx*dx + dy*dy))
			var r, g, b, a uint8
			if gradientPos < 0.5 {
				pos := gradientPos * 2
				r = uint8(float64(startColor.R)*(1-pos) + float64(midColor.R)*pos)
				g = uint8(float64(startColor.G)*(1-pos) + float64(midColor.G)*pos)
				b = uint8(float64(startColor.B)*(1-pos) + float64(midColor.B)*pos)
				a = uint8(float64(startColor.A)*(1-pos) + float64(midColor.A)*pos)
			} else {
				pos := (gradientPos - 0.5) * 2
				r = uint8(float64(midColor.R)*(1-pos) + float64(endColor.R)*pos)
				g = uint8(float64(midColor.G)*(1-pos) + float64(endColor.G)*pos)
				b = uint8(float64(midColor.B)*(1-pos) + float64(endColor.B)*pos)
				a = uint8(float64(midColor.A)*(1-pos) + float64(endColor.A)*pos)
			}

			img.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	return img
}



func main() {
    // Parse arguments
    width := flag.Int("size", 256, "Width of the image")
	height := flag.Int("h", 256, "Height of the image")
    angle := flag.Float64("angle", 0.0, "Angle of rotation in degrees")
    output := flag.String("o", "output.jpg", "Output file name")
    flag.Parse()

    // Draw gradient
    img := drawGradient(*width, *height, *angle, color.RGBA{254,190,90, 255}, color.RGBA{222,50,76, 255},color.RGBA{140,67,174,255})

    // Save image to file
	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("Image saved to %s\n", *output)
}
