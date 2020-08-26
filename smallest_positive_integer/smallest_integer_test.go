package smallest_positive_integer_test

import (
	"efficient-algorithms/smallest_positive_integer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	A := []int{-1, -3}

	ans := smallest_positive_integer.Solution(A)

	assert.EqualValues(t, ans, 1)
}
