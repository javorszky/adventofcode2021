package day20

import (
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

type Image struct {
	image       [][]pixel
	enhancement Enhancement
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
	}
}
