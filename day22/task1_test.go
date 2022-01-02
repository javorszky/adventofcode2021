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
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
			want: instruction{
				XFrom: -10,
				XTo:   10,
				YFrom: -10,
				YTo:   10,
				ZFrom: -10,
				ZTo:   10,
				Flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "returns error for an instruction where x is too low",
			args: args{
				in: instruction{
					XFrom: -90,
					XTo:   -55,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where x is too high",
			args: args{
				in: instruction{
					XFrom: 55,
					XTo:   99,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where y is too low",
			args: args{
				in: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -100,
					YTo:   -55,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where z is too low",
			args: args{
				in: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -100,
					ZTo:   -55,
					Flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where y is too high",
			args: args{
				in: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: 55,
					YTo:   100,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "returns error for an instruction where z is too high",
			args: args{
				in: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 55,
					ZTo:   100,
					Flip:  on,
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

func Test_task1(t *testing.T) {
	type args struct {
		input []instruction
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "task 1 super smol example",
			args: args{input: testInput(t, "example_input_smol.txt")},
			want: 39,
		},
		{
			name: "task 1 example",
			args: args{input: testInput(t, "example_input.txt")},
			want: 590784, // the max that COULD be on it 1,030,301 (101 x 101 x 101)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.args.input), "task1(%v)", tt.args.input)
		})
	}
}

func testInput(t testing.TB, fn string) []instruction {
	t.Helper()

	input := getInputs(fn)
	insts := make([]instruction, len(input))

	for i, line := range input {
		insts[i] = parseInstruction(line)
	}

	return insts
}
