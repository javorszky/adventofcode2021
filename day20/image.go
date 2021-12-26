package day20

import (
	"log"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

var sweep = []int{-1, 0, 1}

type Image struct {
	image       [][]pixel
	enhancement Enhancement
	xMax, yMax  int
	outside     pixel
	allDark     pixel
	allLit      pixel
}

func NewImage(img string, enhance string) Image {
	rows := strings.Split(img, util.NewLine)
	imgMultiSlice := make([][]pixel, len(rows)+2)
	imgMultiSlice[0] = make([]pixel, len(rows[0])+2)
	imgMultiSlice[len(rows)+1] = make([]pixel, len(rows[0])+2)

	for i := 0; i < len(rows[0])+2; i++ {
		imgMultiSlice[0][i] = darkPx
		imgMultiSlice[len(rows)+1][i] = darkPx
	}

	for rowNum, row := range rows {
		for colNum, px := range row {
			if colNum == 0 {
				imgMultiSlice[rowNum+1] = make([]pixel, len(row)+2)
				imgMultiSlice[rowNum+1][0] = darkPx
			}

			imgMultiSlice[rowNum+1][colNum+1] = parseChar(uint8(px))
		}

		imgMultiSlice[rowNum+1][len(imgMultiSlice[rowNum+1])-1] = darkPx
	}

	return Image{
		image:       imgMultiSlice,
		enhancement: NewEnhancement(enhance),
		yMax:        len(rows) + 1,
		xMax:        len(rows[0]) + 1,
		outside:     darkPx,
		allDark:     parseChar(enhance[0]),
		allLit:      parseChar(enhance[len(enhance)-1]),
	}
}

func (i Image) tick() Image {
	newImage := Image{
		image:       nil,
		enhancement: i.enhancement,
		xMax:        i.xMax + 2,
		yMax:        i.yMax + 2,
		allDark:     i.allDark,
		allLit:      i.allLit,
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

					if checkRow < 1 || checkCol < 1 || checkRow > i.xMax+1 || checkCol > i.yMax+1 {
						subPixel[c] = i.outside
						c++

						continue
					}

					subPixel[c] = i.image[checkRow-1][checkCol-1]
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

	switch i.outside {
	case lightPx:
		newImage.outside = i.allLit
	case darkPx:
		newImage.outside = i.allDark
	default:
		log.Fatalf("i.outside is not one of the two possibilities.")
	}

	return newImage
}

func (i Image) Lights() int {
	lights := 0

	for _, row := range i.image {
		for _, col := range row {
			if col == lightPx {
				lights++
			}
		}
	}

	return lights
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
