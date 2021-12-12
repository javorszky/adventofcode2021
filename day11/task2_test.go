package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simulateUntilSync(t *testing.T) {
	tests := []struct {
		name  string
		board func() map[uint]uint
		want  int
	}{
		{
			name: "example simul flash",
			board: func() map[uint]uint {
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
			want: 195,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, simulateUntilSync(tt.board()), "simulateUntilSync(%v)", tt.board())
		})
	}
}
