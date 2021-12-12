package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getNeighbourCoords(t *testing.T) {
	tests := []struct {
		name string
		in   uint
		want []uint
	}{
		{
			name: "gets neighbours for top left corner",
			in:   0,
			want: []uint{
				0b00000001,
				0b00010000,
				0b00010001,
			},
		},
		{
			name: "gets neighbours for top right corner",
			in:   0b00001001,
			want: []uint{
				0b00001000, // left
				0b00011001, // below
				0b00011000, // left and below
			},
		},
		{
			name: "gets neighbours bottom left corner",
			in:   0b10010000,
			want: []uint{
				0b10010001, // right
				0b10000001, // above and right
				0b10000000, // above
			},
		},
		{
			name: "gets neighbours bottom right corner",
			in:   0b10011001,
			want: []uint{
				0b10011000, // left
				0b10001000, // above left
				0b10001001, // above
			},
		},
		{
			name: "gets neighbours bottom non-corner piece",
			in:   0b10010101,
			want: []uint{
				0b10010100, // left
				0b10000100, // above left
				0b10000101, // above
				0b10000110, // above right
				0b10010110, // right
			},
		},
		{
			name: "gets neighbours top non-corner piece",
			in:   0b00000101,
			want: []uint{
				0b00000100, // left
				0b00000110, // bottom right
				0b00010100, // bottom left
				0b00010101, // bottom
				0b00010110, // right
			},
		},
		{
			name: "gets neighbours left non-corner piece",
			in:   0b01010000,
			want: []uint{
				0b01000000, // above
				0b01000001, // above right
				0b01010001, // right
				0b01100001, // bottom right
				0b01100000, // bottom
			},
		},
		{
			name: "gets neighbours right non-corner piece",
			in:   0b01011001,
			want: []uint{
				0b01001001, // above
				0b01001000, // above left
				0b01011000, // left
				0b01101000, // below left
				0b01101001, // below
			},
		},
		{
			name: "gets neighbours mid piece",
			in:   0b01010101,
			want: []uint{
				0b01000100, // above left
				0b01000101, // above
				0b01000110, // above right
				0b01010110, // right
				0b01100110, // below right
				0b01100101, // below
				0b01100100, // below left
				0b01010100, // below
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, getNeighbourCoords(tt.in), "getNeighbourCoords(%v)", tt.in)
		})
	}
}

func Test_step(t *testing.T) {
	tests := []struct {
		name  string
		m     func() map[uint]uint
		want  func() map[uint]uint
		want1 int
	}{
		{
			name: "example init -> step 1",
			m: func() map[uint]uint {
				return parseIntoGrid([]string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				})
			},
			want: func() map[uint]uint {
				return parseIntoGrid([]string{
					"6594254334",
					"3856965822",
					"6375667284",
					"7252447257",
					"7468496589",
					"5278635756",
					"3287952832",
					"7993992245",
					"5957959665",
					"6394862637",
				})
			},
			want1: 0,
		},
		{
			name: "example step 1 -> step 2",
			m: func() map[uint]uint {
				return parseIntoGrid([]string{
					"6594254334",
					"3856965822",
					"6375667284",
					"7252447257",
					"7468496589",
					"5278635756",
					"3287952832",
					"7993992245",
					"5957959665",
					"6394862637",
				})
			},
			want: func() map[uint]uint {
				return parseIntoGrid([]string{
					"8807476555",
					"5089087054",
					"8597889608",
					"8485769600",
					"8700908800",
					"6600088989",
					"6800005943",
					"0000007456",
					"9000000876",
					"8700006848",
				})
			},
			want1: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := step(tt.m())
			assert.Equalf(t, tt.want(), got, "step(%v)", tt.m)
			assert.Equalf(t, tt.want1, got1, "step(%v)", tt.m)
		})
	}
}

func Test_inc1(t *testing.T) {
	tests := []struct {
		name     string
		m        func() map[uint]uint
		want     func() map[uint]uint
		wantLeft func() map[uint]uint
	}{
		{
			name: "increments table once",
			m: func() map[uint]uint {
				return parseIntoGrid([]string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				})
			},
			want: func() map[uint]uint {
				return parseIntoGrid([]string{
					"6594254334",
					"3856965822",
					"6375667284",
					"7252447257",
					"7468496589",
					"5278635756",
					"3287952832",
					"7993992245",
					"5957959665",
					"6394862637",
				})
			},
			wantLeft: func() map[uint]uint {
				return parseIntoGrid([]string{
					"6594254334",
					"3856965822",
					"6375667284",
					"7252447257",
					"7468496589",
					"5278635756",
					"3287952832",
					"7993992245",
					"5957959665",
					"6394862637",
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotLeft := inc1(tt.m())
			assert.Equalf(t, tt.want(), got, "inc1(%v) - got", tt.m())
			assert.Equalf(t, tt.wantLeft(), gotLeft, "inc1(%v) - gotLeft", tt.m())
		})
	}
}

func Test_getSum(t *testing.T) {
	tests := []struct {
		name string
		m    func() map[uint]uint
		want string
	}{
		{
			name: "Gets state of board",
			m: func() map[uint]uint {
				return parseIntoGrid([]string{
					"6594254334",
					"3856965822",
					"6375667284",
					"7252447257",
					"7468496589",
					"5278635756",
					"3287952832",
					"7993992245",
					"5957959665",
					"6394862637",
				})
			},
			want: "6594254334385696582263756672847252447257746849658952786357563287952832799399224559579596656394862637",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getSum(tt.m()), "getSum(%v)", tt.m())
		})
	}
}
