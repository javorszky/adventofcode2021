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
