package alphabet_symmetry_test

import (
	"efficient-algorithms/alphabet_symmetry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {

	assert.Equal(t, alphabet_symmetry.Solve2([]string{"abcde", "gBa"}), []int{5, 1})
}
