package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bingoBoard_Mark(t *testing.T) {
	defaultPieces := map[int]uint{
		67: 0b1000000000000000000000000,
		97: 0b0100000000000000000000000,
		50: 0b0010000000000000000000000,
		51: 0b0001000000000000000000000,
		1:  0b0000100000000000000000000,
		47: 0b0000010000000000000000000,
		15: 0b0000001000000000000000000,
		77: 0b0000000100000000000000000,
		31: 0b0000000010000000000000000,
		66: 0b0000000001000000000000000,
		24: 0b0000000000100000000000000,
		14: 0b0000000000010000000000000,
		55: 0b0000000000001000000000000,
		70: 0b0000000000000100000000000,
		52: 0b0000000000000010000000000,
		76: 0b0000000000000001000000000,
		46: 0b0000000000000000100000000,
		19: 0b0000000000000000010000000,
		32: 0b0000000000000000001000000,
		73: 0b0000000000000000000100000,
		34: 0b0000000000000000000010000,
		22: 0b0000000000000000000001000,
		54: 0b0000000000000000000000100,
		75: 0b0000000000000000000000010,
		17: 0b0000000000000000000000001,
	}

	tests := []struct {
		name           string
		board          bingoBoard
		mark           int
		want           bool
		wantBoardState uint
	}{
		{
			name: "marks a piece correctly",
			board: bingoBoard{
				pieces: defaultPieces,
				state:  emptyState,
			},
			mark:           46,
			want:           false,
			wantBoardState: 0b0000000000000000100000000,
		},
		{
			name: "handles marking a piece that's not on the board correctly",
			board: bingoBoard{
				pieces: defaultPieces,
				state:  emptyState,
			},
			mark:           84,
			want:           false,
			wantBoardState: emptyState,
		},
		{
			name: "handles marking a new number that results in a win",
			board: bingoBoard{
				pieces: defaultPieces,
				state:  0b1011001010100111010101011,
			},
			mark:           32,
			want:           true,
			wantBoardState: 0b1011001010100111011101011,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.board.Mark(tt.mark), "Mark(%v)", tt.mark)
			assert.Equal(t, tt.wantBoardState, tt.board.state)
		})
	}
}
