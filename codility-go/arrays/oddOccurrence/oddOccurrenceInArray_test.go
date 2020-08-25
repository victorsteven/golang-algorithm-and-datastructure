package oddOccurrence_test

import (
	"efficient-algorithms/codility-go/arrays/oddOccurrence"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddOccurrence(t *testing.T) {

	A := []int{9, 1, 9, 3, 9, 7, 9}

	num := oddOccurrence.OddOccurrence(A)

	assert.EqualValues(t, num, 7)
}
