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

func drawGradient(width, height int, angle float64, colors ...color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{colors[0]}, image.ZP, draw.Src)

    var isFlipped bool
    if int(angle) / 90 % 2 != 0 {
        isFlipped = true
        angle = angle - 90
    }

    // Calculate the direction of the gradient
	dx := float64(width) * math.Sin(angle*math.Pi/180)
	dy := float64(height) * math.Cos(angle*math.Pi/180)

	// Define the number of segments based on the number of colors
	numSegments := len(colors) - 1
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// Calculate the position of the current pixel along the gradient
			gradientPos := (float64(i)*dx + float64(j)*dy) / (float64(width)*dx + float64(height)*dy)

			// Find the two colors that this pixel lies between
			segment := int(gradientPos * float64(numSegments))
			startColor := colors[segment]
			endColor := colors[segment+1]

			// Calculate the color of the current pixel by interpolating between the two colors
			segmentPos := gradientPos*float64(numSegments) - float64(segment)
			r := uint8(float64(startColor.R)*(1-segmentPos) + float64(endColor.R)*segmentPos)
			g := uint8(float64(startColor.G)*(1-segmentPos) + float64(endColor.G)*segmentPos)
			b := uint8(float64(startColor.B)*(1-segmentPos) + float64(endColor.B)*segmentPos)
			a := uint8(float64(startColor.A)*(1-segmentPos) + float64(endColor.A)*segmentPos)
			img.Set(i, j, color.RGBA{r, g, b, a})
		}
	}

    if isFlipped {
        img = flipHorizontal(img)
    }

	return img
}

func flipHorizontal(img *image.RGBA) *image.RGBA {
    // Copy the original image
    flipped := image.NewRGBA(img.Bounds())
    draw.Draw(flipped, flipped.Bounds(), img, image.Point{}, draw.Src)

    // Flip each row of pixels horizontally
    for y := 0; y < img.Bounds().Dy(); y++ {
        for x1, x2 := 0, img.Bounds().Dx()-1; x1 < x2; x1, x2 = x1+1, x2-1 {
            // Swap the pixels on either side of the row's midpoint
            flipped.Set(x1, y, img.At(x2, y))
            flipped.Set(x2, y, img.At(x1, y))
        }
    }

    return flipped
}

// Convert hex color code to RGBA color
func hexToRGBA(hexCode string) (color.RGBA, error) {
	if len(hexCode) != 6 {
		return color.RGBA{}, fmt.Errorf("Invalid color code length: %s", hexCode)
	}
	var r, g, b uint8
	_, err := fmt.Sscanf(hexCode, "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{r, g, b, 255}, nil
}

func main() {
    // Parse arguments
    width := flag.Int("w", 256, "Width of the image")
	height := flag.Int("h", 256, "Height of the image")
    angle := flag.Float64("angle", 0.0, "Angle of rotation in degrees")
    output := flag.String("o", "output.jpg", "Output file name")
	flag.Parse()

	colorArgs := flag.Args()
	if len(colorArgs) == 0 {
		log.Fatal("At least one color code must be provided")
	}

	// Convert color codes to RGBA colors
	var colors []color.RGBA
	for _, code := range colorArgs {
		c, err := hexToRGBA(code)
		if err != nil {
			log.Fatalf("Invalid color code: %s", code)
		}
		colors = append(colors, c)
	}

	// Draw gradient
	img := drawGradient(*width, *height, *angle, colors...)

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
