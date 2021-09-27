package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestNumber(t *testing.T) {
	assert.Equal(t, 5, smallest_number([]int{6,7,5,8,9,10,13,100}))
}

func TestLargestNumber(t *testing.T) {
	assert.Equal(t, 100, largest_number([]int{6,7,5,8,9,10,13,100}))
}
