package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// 画像サイズを設定
	width := 256
	height := 256

	// 画像を生成
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// グラデーションの色を設定
	color1 := color.RGBA{255, 0, 0, 255} // 赤
	color2 := color.RGBA{0, 0, 255, 255} // 青

	// グラデーションを生成
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ratio := float64(x) / float64(width)
			r := uint8(float64(color1.R)*(1-ratio) + float64(color2.R)*ratio)
			g := uint8(float64(color1.G)*(1-ratio) + float64(color2.G)*ratio)
			b := uint8(float64(color1.B)*(1-ratio) + float64(color2.B)*ratio)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// 画像をJPEG形式で保存
	file, err := os.Create("output/gradient.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{100})
}
