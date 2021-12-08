package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  interface{}
	}{
		{
			name:  "example input done right",
			input: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, task1(tt.input))
		})
	}
}
