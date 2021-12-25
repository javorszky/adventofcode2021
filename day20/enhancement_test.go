package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnhancement(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want Enhancement
	}{
		{
			name: "creates new enhanement map",
			args: args{s: "...#...##.####....#.#####.......#.#"},
			want: Enhancement{s: "...#...##.####....#.#####.......#.#"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEnhancement(tt.args.s))
		})
	}
}

func TestEnhancement_subPixel(t *testing.T) {
	type fields struct {
		s string
	}

	type args struct {
		in [9]pixel
	}

	enhancementMap := "##..##...#..#.#...#.....##....#.#..#.#.####...#####..#.######...#......#.##.#..######.########" +
		"....#.#..##.####.##...##..#.########.#.##..........##.######.#......#..#...##..#..#.###.#.#..#..#...##.###.." +
		"...#.#.###..##.####....##.#....#.#.###..###.....####..###..##.#.#..##....#.#....#####.##.....#.#.#..###..#.." +
		"...####..##.#..#.###....#...##..###.#.###.##.####..#.##......##.##.#.##..##...##..######..####.#.##..###..##" +
		"#.###.##.##..###..#.......##.#######.#..##..##.###.#.#.#.####...####..#.#.#.......##.##.#....."

	tests := []struct {
		name   string
		fields fields
		args   args
		want   pixel
	}{
		{
			name:   "using actual input enhancement map, correctly returns pixel by coordinate 13",
			fields: fields{s: enhancementMap},
			args:   args{in: [9]pixel{darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, lightPx}},
			want:   darkPx,
		},
		{
			name:   "using actual input enhancement map, correctly returns pixel by coordinate 12",
			fields: fields{s: enhancementMap},
			args:   args{in: [9]pixel{darkPx, darkPx, darkPx, darkPx, darkPx, lightPx, lightPx, darkPx, darkPx}},
			want:   lightPx,
		},
		{
			name:   "using actual input enhancement map, correctly returns the very first pixel",
			fields: fields{s: enhancementMap},
			args:   args{in: [9]pixel{darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx, darkPx}},
			want:   lightPx,
		},
		{
			name:   "using actual input enhancement map, correctly returns the very last pixel",
			fields: fields{s: enhancementMap},
			args:   args{in: [9]pixel{lightPx, lightPx, lightPx, lightPx, lightPx, lightPx, lightPx, lightPx, lightPx}},
			want:   darkPx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnhancement(tt.fields.s)
			assert.Equalf(t, tt.want, e.subPixel(tt.args.in), "subPixel(%v)", tt.args.in)
		})
	}
}
