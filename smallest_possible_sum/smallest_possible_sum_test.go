package smallest_possible_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	assert.Equal(t, Solution([]int{6, 9, 12}), 9)
}
