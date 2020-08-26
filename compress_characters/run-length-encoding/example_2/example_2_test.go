package example_2_test

import (
	"efficient-algorithms/compress_characters/run-length-encoding/example_2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {

	sample := []struct{
		str string
		result string
	}{
		{
			"AAAADDFFFAAFFDDSSWW",
			"4A2D3F2A2F2D2S2W",
		},
		{
			"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW",
			"12WB12W3B24WB14W",
		},
		{
			"ABC",
			"ABC",
		},
	}

	for _, v := range sample {

		result := example_2.Solution1(v.str)

		assert.EqualValues(t, v.result, result)
	}
}


func TestSolution2(t *testing.T) {

	sample := []struct{
		str string
		result string
	}{
		{
			"AAAADDFFFAAFFDDSSWW",
			"4A2D3F2A2F2D2S2W",
		},
		{
			"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW",
			"12WB12W3B24WB14W",
		},
		{
			"ABC",
			"ABC",
		},
	}

	for _, v := range sample {

		result := example_2.Solution2(v.str)

		assert.EqualValues(t, v.result, result)
	}
}


