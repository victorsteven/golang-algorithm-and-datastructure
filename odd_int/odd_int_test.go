package odd_int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOdd(t *testing.T) {

	assert.Equal(t, FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}), 5)
	assert.Equal(t, FindOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}), -1)
	assert.Equal(t, FindOdd([]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}), 5)
	assert.Equal(t, FindOdd([]int{10}), 10)
	assert.Equal(t, FindOdd([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}), 10)
	assert.Equal(t, FindOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}), 1)
}
