package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fetchCmdAndValue(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		want  uint8
		want1 uint8
	}{
		{
			name:  "fetches correct instructions for up 1",
			line:  "up 1",
			want:  0x75,
			want1: 0x31,
		},
		{
			name:  "fetches correct instructions for down 2",
			line:  "down 2",
			want:  0x64,
			want1: 0x32,
		},
		{
			name:  "fetches correct instructions for forward 3",
			line:  "forward 3",
			want:  0x66,
			want1: 0x33,
		},
		{
			name:  "fetches correct instructions for up 4",
			line:  "up 4",
			want:  0x75,
			want1: 0x34,
		},
		{
			name:  "fetches correct instructions for up 5",
			line:  "up 5",
			want:  0x75,
			want1: 0x35,
		},
		{
			name:  "fetches correct instructions for up 6",
			line:  "up 6",
			want:  0x75,
			want1: 0x36,
		}, {
			name:  "fetches correct instructions for up 7",
			line:  "up 7",
			want:  0x75,
			want1: 0x37,
		},
		{
			name:  "fetches correct instructions for up 8",
			line:  "up 8",
			want:  0x75,
			want1: 0x38,
		},
		{
			name:  "fetches correct instructions for up 9",
			line:  "up 9",
			want:  0x75,
			want1: 0x39,
		},
		{
			name:  "fetches correct instructions for up 0",
			line:  "up 0",
			want:  0x75,
			want1: 0x30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fetchCmdAndValue(tt.line)
			assert.Equal(t, tt.want, got, "wrong code for instruction")
			assert.Equal(t, tt.want1, got1, "wrong code for value")
		})
	}
}
