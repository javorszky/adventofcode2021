package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	type args struct {
		player1 int
		player2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example case",
			args: args{
				player1: 4,
				player2: 8,
			},
			want: 739785,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.args.player1, tt.args.player2),
				"task1(%v, %v)", tt.args.player1, tt.args.player2)
		})
	}
}
