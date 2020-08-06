// 5 * Напишите программу, генерирующую png-файл с рисунком шахматной доски.

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

	boardNum := 240

	myimage := image.NewRGBA(image.Rect(0, 0, boardNum, boardNum))
	colors := make(map[int]color.RGBA, 2)

	colors[0] = color.RGBA{3, 2, 71, 255}    // синий тем
	colors[1] = color.RGBA{3, 138, 168, 255} // синий светл

	numColor := 0
	sizeBoard := 8
	sizeBlock := int(boardNum / sizeBoard)
	loc_x := 0

	// эту часть где-то подсмотрел
	for curr_x := 0; curr_x < sizeBoard; curr_x++ {

		loc_y := 0
		for curr_y := 0; curr_y < sizeBoard; curr_y++ {

			draw.Draw(myimage, image.Rect(loc_x, loc_y, loc_x+sizeBlock, loc_y+sizeBlock),
				&image.Uniform{colors[numColor]}, image.ZP, draw.Src)

			loc_y += sizeBlock
			numColor = 1 - numColor
		}
		loc_x += sizeBlock
		numColor = 1 - numColor
	}

	file, err := os.Create("chess.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, myimage)

}
