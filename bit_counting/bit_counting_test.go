package bit_counting_test

import (
	"efficient-algorithms/bit_counting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBits(t *testing.T) {

	assert.Equal(t, bit_counting.CountBits(1234), 5)
}
