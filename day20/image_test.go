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
				xMax:        4,
				yMax:        4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewImage(tt.args.img, tt.args.enhance), "NewImage(%v, %v)", tt.args.img, tt.args.enhance)
		})
	}
}

func TestImage_tick(t *testing.T) {
	enhance := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..##" +
		"###..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##." +
		".....#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#..." +
		"##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#..." +
		"...#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
	img := Image{
		image: [][]pixel{
			{lightPx, darkPx, darkPx, lightPx, darkPx},
			{lightPx, darkPx, darkPx, darkPx, darkPx},
			{lightPx, lightPx, darkPx, darkPx, lightPx},
			{darkPx, darkPx, lightPx, darkPx, darkPx},
			{darkPx, darkPx, lightPx, lightPx, lightPx},
		},
		enhancement: Enhancement{s: enhance},
		xMax:        4,
		yMax:        4,
	}

	tests := []struct {
		name  string
		img   Image
		ticks int
		want  Image
	}{
		{
			name:  "ticks example start once",
			img:   img,
			ticks: 1,
			want: Image{
				image: [][]pixel{
					{darkPx, lightPx, lightPx, darkPx, lightPx, lightPx, darkPx},
					{lightPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx},
					{lightPx, lightPx, darkPx, lightPx, darkPx, darkPx, lightPx},
					{lightPx, lightPx, lightPx, lightPx, darkPx, darkPx, lightPx},
					{darkPx, lightPx, darkPx, darkPx, lightPx, lightPx, darkPx},
					{darkPx, darkPx, lightPx, lightPx, darkPx, darkPx, lightPx},
					{darkPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx},
				},
				enhancement: Enhancement{s: enhance},
				xMax:        6,
				yMax:        6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localImg := tt.img
			for i := 0; i < tt.ticks; i++ {
				localImg = tt.img.tick()
			}

			assert.Equalf(t, tt.want, localImg, "tick()")
		})
	}
}

func TestImage_String(t *testing.T) {
	enhance := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..##" +
		"###..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##." +
		".....#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#..." +
		"##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#..." +
		"...#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
	img := Image{
		image: [][]pixel{
			{lightPx, darkPx, darkPx, lightPx, darkPx},
			{lightPx, darkPx, darkPx, darkPx, darkPx},
			{lightPx, lightPx, darkPx, darkPx, lightPx},
			{darkPx, darkPx, lightPx, darkPx, darkPx},
			{darkPx, darkPx, lightPx, lightPx, lightPx},
		},
		enhancement: Enhancement{s: enhance},
		xMax:        4,
		yMax:        4,
	}

	tests := []struct {
		name string
		img  Image
		want string
	}{
		{
			name: "image.String() works on example",
			img:  img,
			want: `#..#.
#....
##..#
..#..
..###`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.img.String(), "String()")
		})
	}
}
