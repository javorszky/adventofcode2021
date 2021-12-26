package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_assembledTask1(t *testing.T) {
	n := assembledTask1()

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, n.value())
		n = n.next()
	}

	assert.Equal(t, 1, n.value())

	for i := 10; i >= 1; i-- {
		n = n.previous()
		assert.Equal(t, i, n.value())
	}
}

func Test_node_rotateTo(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want int
	}{
		{
			name: "rotates to 5",
			v:    5,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := assembledTask1()
			assert.Equalf(t, tt.want, n.rotateTo(tt.v).value(), "rotateTo(%v)", tt.v)
		})
	}
}

func Test_node_rotateBy(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want int
	}{
		{
			name: "rotates starter ring by 4",
			v:    4,
			want: 5,
		},
		{
			name: "rotates starter ring by 10",
			v:    10,
			want: 1,
		},
		{
			name: "rotates starter ring by 205",
			v:    205,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := assembledTask1()

			got, got1 := n.rotateBy(tt.v)
			assert.Equalf(t, tt.want, got.value(), "rotateBy(%v)", tt.v)
			assert.Equalf(t, tt.want, got1, "rotateBy(%v)", tt.v)
		})
	}
}
