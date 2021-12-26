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
					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, lightPx, darkPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, darkPx, darkPx, darkPx},
					{darkPx, lightPx, lightPx, darkPx, darkPx, lightPx, darkPx},
					{darkPx, darkPx, darkPx, lightPx, darkPx, darkPx, darkPx},
					{darkPx, darkPx, darkPx, lightPx, lightPx, lightPx, darkPx},
					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
				},
				enhancement: Enhancement{s: enhance},
				xMax:        6,
				yMax:        6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newImage := NewImage(tt.args.img, tt.args.enhance)
			assert.Equalf(t, tt.want, newImage, "expected:\n\n%s\n\nactual:\n\n%s\n\n", tt.want.String(), newImage.String())
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
			{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
			{darkPx, lightPx, darkPx, darkPx, lightPx, darkPx, darkPx},
			{darkPx, lightPx, darkPx, darkPx, darkPx, darkPx, darkPx},
			{darkPx, lightPx, lightPx, darkPx, darkPx, lightPx, darkPx},
			{darkPx, darkPx, darkPx, lightPx, darkPx, darkPx, darkPx},
			{darkPx, darkPx, darkPx, lightPx, lightPx, lightPx, darkPx},
			{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
		},
		enhancement: Enhancement{s: enhance},
		xMax:        6,
		yMax:        6,
		outside:     darkPx,
		allDark:     darkPx,
		allLit:      lightPx,
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
					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},

					{darkPx, darkPx, lightPx, lightPx, darkPx, lightPx, lightPx, darkPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx, darkPx},
					{darkPx, lightPx, lightPx, darkPx, lightPx, darkPx, darkPx, lightPx, darkPx},
					{darkPx, lightPx, lightPx, lightPx, lightPx, darkPx, darkPx, lightPx, darkPx},
					{darkPx, darkPx, lightPx, darkPx, darkPx, lightPx, lightPx, darkPx, darkPx},
					{darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, darkPx, lightPx, darkPx},
					{darkPx, darkPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx, darkPx},

					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
				},
				enhancement: Enhancement{s: enhance},
				xMax:        8,
				yMax:        8,
				outside:     darkPx,
				allDark:     darkPx,
				allLit:      lightPx,
			},
		},
		{
			name:  "ticks example start twice",
			img:   img,
			ticks: 2,
			want: Image{
				image: [][]pixel{
					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},

					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, darkPx, darkPx},
					{darkPx, darkPx, lightPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx, darkPx, darkPx},
					{darkPx, lightPx, darkPx, lightPx, darkPx, darkPx, darkPx, lightPx, lightPx, lightPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, lightPx, darkPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx},
					{darkPx, darkPx, lightPx, darkPx, lightPx, lightPx, lightPx, lightPx, lightPx, darkPx, darkPx},
					{darkPx, darkPx, darkPx, lightPx, darkPx, lightPx, lightPx, lightPx, lightPx, lightPx, darkPx},
					{darkPx, darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, lightPx, lightPx, darkPx, darkPx},
					{darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, lightPx, lightPx, darkPx, darkPx, darkPx},

					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx},
				},
				enhancement: Enhancement{s: enhance},
				xMax:        10,
				yMax:        10,
				outside:     darkPx,
				allDark:     darkPx,
				allLit:      lightPx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localImg := tt.img
			for i := 0; i < tt.ticks; i++ {
				localImg = localImg.tick()
			}

			assert.Equalf(t, tt.want, localImg,
				"want:\n\n%s\n\nactual:\n\n%s\n\n", tt.want.String(), localImg.String())
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

func TestImage_Lights(t *testing.T) {
	enhance := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..##" +
		"###..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##." +
		".....#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#..." +
		"##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#..." +
		"...#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"

	tests := []struct {
		name string
		img  Image
		want int
	}{
		{
			name: "counts lights correctly",
			img: Image{
				image: [][]pixel{
					{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, darkPx},
					{darkPx, lightPx, darkPx, darkPx, lightPx, darkPx, lightPx, darkPx, darkPx},
					{lightPx, darkPx, lightPx, darkPx, darkPx, darkPx, lightPx, lightPx, lightPx},
					{lightPx, darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, lightPx, darkPx},
					{lightPx, darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, darkPx, lightPx},
					{darkPx, lightPx, darkPx, lightPx, lightPx, lightPx, lightPx, lightPx, darkPx},
					{darkPx, darkPx, lightPx, darkPx, lightPx, lightPx, lightPx, lightPx, lightPx},
					{darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, lightPx, lightPx, darkPx},
					{darkPx, darkPx, darkPx, darkPx, lightPx, lightPx, lightPx, darkPx, darkPx},
				},
				enhancement: Enhancement{s: enhance},
				xMax:        8,
				yMax:        8,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.img.Lights(), "Lights()")
		})
	}
}
