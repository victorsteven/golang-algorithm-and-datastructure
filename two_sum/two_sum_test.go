package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	assert.Equal(t, TwoSum([]int{1, 2, 3}, 4), [2]int{0, 2})

	assert.Equal(t, TwoSum2([]int{1234, 5678, 9012}, 14690), [2]int{1, 2})

}
