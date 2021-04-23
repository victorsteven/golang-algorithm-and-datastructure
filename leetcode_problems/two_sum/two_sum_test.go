package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	assert.Equal(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, TwoSum([]int{3, 2, 4}, 6))

	assert.Equal(t, []int{0, 1}, TwoSum2([]int{2, 7, 11, 15}, 9))

	assert.Equal(t, []int{0, 1}, TwoSum2a([]int{2, 7, 11, 15}, 9))

	assert.Equal(t, []int{1, 2}, TwoSum2([]int{3, 2, 4}, 6))

	assert.Equal(t, []int{0, 1}, TwoSum3([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, TwoSum3([]int{3, 2, 4}, 6))

	assert.Equal(t, []int{1, 2}, TwoSum3a([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{1, 2}, TwoSum3b([]int{3, 2, 4}, 6))

}
