package example_1_test

import (
	"efficient-algorithms/compress_characters/run-length-encoding/example_1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {

	sample := []struct {
		str    string
		result string
	}{
		{
			"AAAADDFFFAAFFDDSSWW",
			"4A2D3F2A2F2D2S2W",
		},
		{
			"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW",
			"12W1B12W3B24W1B14W",
		},
		{
			"ABC",
			"1A1B1C",
		},
	}

	for _, v := range sample {

		result := example_1.Solution1(v.str)

		assert.EqualValues(t, v.result, result)
	}
}

func TestSolution2(t *testing.T) {

	sample := []struct {
		str    string
		result string
	}{
		{
			"AAAADDFFFAAFFDDSSWW",
			"4A2D3F2A2F2D2S2W",
		},
		{
			"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW",
			"12W1B12W3B24W1B14W",
		},
		{
			"ABC",
			"1A1B1C",
		},
	}

	for _, v := range sample {

		result := example_1.Solution2(v.str)

		assert.EqualValues(t, v.result, result)
	}
}
