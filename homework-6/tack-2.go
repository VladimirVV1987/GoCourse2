// 2 Дополните пример из раздела «Пакет img» изображением
//горизонтальных и вертикальных линий. Воспользуйтесь
//статьей «Работа с изображениями».
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	// линия горизонт

	var red color.Color = color.RGBA{200, 30, 30, 255}

	for x := 5; x < 180; x++ {
		y := x/200 + 100
		rectImg.Set(x, y, red)
	}

	// линия вертикаль

	var blue color.Color = color.RGBA{0, 0, 200, 255}

	for y := 5; y < 180; y++ {
		x := y/200 + 100
		rectImg.Set(x, y, blue)
	}

	file, err := os.Create("t2-rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
