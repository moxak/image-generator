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

func drawGradient(width, height int, angle float64, startColor, endColor color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{startColor}, image.ZP, draw.Src)

	dx := float64(width) * math.Cos(angle*math.Pi/180)
	dy := float64(height) * math.Sin(angle*math.Pi/180)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// calculate linear interpolation between startColor and endColor
			// based on position in the gradient
			gradientPos := ((float64(x)*dx + float64(y)*dy) / (dx*dx + dy*dy))
			r := uint8(float64(startColor.R)*(1-gradientPos) + float64(endColor.R)*gradientPos)
			g := uint8(float64(startColor.G)*(1-gradientPos) + float64(endColor.G)*gradientPos)
			b := uint8(float64(startColor.B)*(1-gradientPos) + float64(endColor.B)*gradientPos)
			a := uint8(float64(startColor.A)*(1-gradientPos) + float64(endColor.A)*gradientPos)

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
    img := drawGradient(*width, *height, *angle, color.RGBA{254,190,90, 255}, color.RGBA{140,67,174,255})

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
