package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeInstructionMap(t *testing.T) {
	type args struct {
		overlaps map[string]instruction
	}

	tests := []struct {
		name string
		args args
		want map[string]instruction
	}{
		{
			name: "merges instructions",
			args: args{
				overlaps: map[string]instruction{
					// box back face
					"-20/-6/-5/5/-5/5/on": {
						XFrom: -20,
						XTo:   -6,
						YFrom: -5,
						YTo:   5,
						ZFrom: -5,
						ZTo:   5,
						Flip:  on,
					},
					// box bottom face
					"-5/5/-5/5/-20/-6/on": {
						XFrom: -5,
						XTo:   5,
						YFrom: -5,
						YTo:   5,
						ZFrom: -20,
						ZTo:   -6,
						Flip:  on,
					},
					// box left face
					"-5/5/-20/-6/-5/5/on": {
						XFrom: -5,
						XTo:   5,
						YFrom: -20,
						YTo:   -6,
						ZFrom: -5,
						ZTo:   5,
						Flip:  on,
					},
					// box back left edge
					"-20/-6/-20/-6/-5/5/on": {
						XFrom: -20,
						XTo:   -6,
						YFrom: -20,
						YTo:   -6,
						ZFrom: -5,
						ZTo:   5,
						Flip:  on,
					},
					// box back bottom edge
					"-20/-6/-5/5/-20/-6/on": {
						XFrom: -20,
						XTo:   -6,
						YFrom: -5,
						YTo:   5,
						ZFrom: -20,
						ZTo:   -6,
						Flip:  on,
					},
					// box left bottom edge
					"-5/5/-20/-6/-20/-6/on": {
						XFrom: -5,
						XTo:   5,
						YFrom: -20,
						YTo:   -6,
						ZFrom: -20,
						ZTo:   -6,
						Flip:  on,
					},
					// box bottom back left corner
					"-20/-6/-20/-6/-20/-6/on": {
						XFrom: -20,
						XTo:   -6,
						YFrom: -20,
						YTo:   -6,
						ZFrom: -20,
						ZTo:   -6,
						Flip:  on,
					},
				},
			},
			want: map[string]instruction{
				// box back face
				// box back left edge
				"-20/-6/-20/5/-5/5/on": {XFrom: -20, XTo: -6, YFrom: -20, YTo: 5, ZFrom: -5, ZTo: 5, Flip: 1},

				// box bottom face
				// box back bottom edge
				// box left bottom edge
				// box bottom back left corner
				"-20/5/-20/5/-20/-6/on": {XFrom: -20, XTo: 5, YFrom: -20, YTo: 5, ZFrom: -20, ZTo: -6, Flip: 1},

				// box left face
				"-5/5/-20/-6/-5/5/on": {XFrom: -5, XTo: 5, YFrom: -20, YTo: -6, ZFrom: -5, ZTo: 5, Flip: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeInstructionMap(tt.args.overlaps), "mergeInstructionMap(%v)", tt.args.overlaps)
		})
	}
}
