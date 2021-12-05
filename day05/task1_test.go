package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterToEuclideanLines(t *testing.T) {
	tests := []struct {
		name   string
		tuples []tuple
		want   []tuple
	}{
		{
			name: "keeps all elements in",
			tuples: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
			want: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
		},
		{
			name:   "handles empty list correctly",
			tuples: []tuple{},
			want:   []tuple{},
		},
		{
			name:   "handles nil list correctly",
			tuples: nil,
			want:   []tuple{},
		},
		{
			name: "removes element because neither dimensions match",
			tuples: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 4},
					{5, 8},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
			want: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterToEuclideanLines(tt.tuples)
			assert.Equal(t, tt.want, got)
		})
	}
}
