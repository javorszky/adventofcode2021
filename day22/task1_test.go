package day22

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_limitInstructionTo50(t *testing.T) {
	type args struct {
		in instruction
	}

	tests := []struct {
		name    string
		args    args
		want    instruction
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "does not touch instruction that's fully within the boundaries",
			args: args{
				in: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
			want: instruction{
				xFrom: -10,
				xTo:   10,
				yFrom: -10,
				yTo:   10,
				zFrom: -10,
				zTo:   10,
				flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "returns error for an instruction where x is too low",
			args: args{
				in: instruction{
					xFrom: -90,
					xTo:   -55,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where x is too high",
			args: args{
				in: instruction{
					xFrom: 55,
					xTo:   99,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where y is too low",
			args: args{
				in: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -100,
					yTo:   -55,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where z is too low",
			args: args{
				in: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -100,
					zTo:   -55,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where y is too high",
			args: args{
				in: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: 55,
					yTo:   100,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where z is too high",
			args: args{
				in: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: 55,
					zTo:   100,
					flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := limitInstructionTo50(tt.args.in)
			if !tt.wantErr(t, err, fmt.Sprintf("limitInstructionTo50(%v)", tt.args.in)) {
				return
			}
			assert.Equalf(t, tt.want, got, "limitInstructionTo50(%v)", tt.args.in)
		})
	}
}
