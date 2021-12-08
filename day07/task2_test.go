package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fuelForDistance(t *testing.T) {
	tests := []struct {
		name     string
		distance int
		want     int
	}{
		{
			name:     "100",
			distance: 100,
			want:     5050,
		},
		{
			name:     "5",
			distance: 5,
			want:     15,
		},
		{
			name:     "6",
			distance: 6,
			want:     21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fuelForDistance(tt.distance), "fuelForDistance(%v)", tt.distance)
		})
	}
}
