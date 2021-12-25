package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	enhance := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..##" +
		"###..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##." +
		".....#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#..." +
		"##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#..." +
		"...#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
	img := `#..#.
#....
##..#
..#..
..###`

	type args struct {
		img     string
		enhance string
	}

	tests := []struct {
		name string
		args args
		want Image
	}{
		{
			name: "parses input strings into image correctly",
			args: args{
				img:     img,
				enhance: enhance,
			},
			want: Image{
				image: [][]pixel{
					{lightPx, darkPx, darkPx, lightPx, darkPx},
					{lightPx, darkPx, darkPx, darkPx, darkPx},
					{lightPx, lightPx, darkPx, darkPx, lightPx},
					{darkPx, darkPx, lightPx, darkPx, darkPx},
					{darkPx, darkPx, lightPx, lightPx, lightPx},
				},
				enhancement: Enhancement{s: enhance},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewImage(tt.args.img, tt.args.enhance), "NewImage(%v, %v)", tt.args.img, tt.args.enhance)
		})
	}
}
