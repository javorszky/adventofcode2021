package day20

import (
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

var sweep = []int{-1, 0, 1}

type Image struct {
	image       [][]pixel
	enhancement Enhancement
	xMax, yMax  int
}

func NewImage(img string, enhance string) Image {
	rows := strings.Split(img, util.NewLine)
	imgMultiSlice := make([][]pixel, len(rows))

	for rowNum, row := range rows {
		for colNum, px := range row {
			if colNum == 0 {
				imgMultiSlice[rowNum] = make([]pixel, len(row))
			}

			imgMultiSlice[rowNum][colNum] = parseChar(uint8(px))
		}
	}

	return Image{
		image:       imgMultiSlice,
		enhancement: NewEnhancement(enhance),
		yMax:        len(rows) - 1,
		xMax:        len(rows[0]) - 1,
	}
}

func (i Image) tick() Image {
	newImage := Image{
		image:       nil,
		enhancement: i.enhancement,
		xMax:        i.xMax + 2,
		yMax:        i.yMax + 2,
	}

	subPixel := [9]pixel{}
	imgMatrix := make([][]pixel, i.xMax+3)

	for row := 0; row <= newImage.xMax; row++ {
		for col := 0; col <= newImage.yMax; col++ {
			c := 0

			for _, deltaRow := range sweep {
				for _, deltaCol := range sweep {
					checkRow := row + deltaRow
					checkCol := col + deltaCol

					if checkRow < 1 || checkCol < 1 || checkRow > i.xMax || checkCol > i.yMax {
						subPixel[c] = darkPx
						c++

						continue
					}

					subPixel[c] = i.image[checkRow][checkCol]
					c++
				}
			}

			if col == 0 {
				imgMatrix[row] = make([]pixel, i.yMax+3)
			}

			imgMatrix[row][col] = newImage.enhancement.subPixel(subPixel)
		}
	}

	newImage.image = imgMatrix

	return newImage
}

func (i Image) String() string {
	sb := strings.Builder{}

	for _, row := range i.image {
		for _, ch := range row {
			sb.WriteString(string(ch))
		}

		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}
