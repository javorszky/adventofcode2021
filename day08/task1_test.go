package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniqueDisplayed(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "no unique in displayed",
			in:   "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | abcdg abcdcg fdcagb acbfg",
			want: 0,
		},
		{
			name: "finds unique in ex 1",
			in:   "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniqueDisplayed(tt.in))
		})
	}
}
