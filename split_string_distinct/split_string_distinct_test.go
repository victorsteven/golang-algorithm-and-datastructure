package split_string_distinct_test

import (
	"efficient-algorithms/split_string_distinct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	samples := []struct {
		s string
		k int
	}{
		{
			"acbbcc",
			6,
		},
		{
			"abb",
			2,
		},
	}

	for _, v := range samples {
		result := split_string_distinct.Solution(v.s)

		assert.EqualValues(t, result, v.k)

	}

}
