package example_3_test

import (
	"efficient-algorithms/compress_characters/run-length-encoding/example_3"
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {

	//sample := []struct{
	//	str string
	//	result string
	//}{
		//6A4D5F2S2W is just one arrangement, so we cant assert for it alone.
		//{
		//	"AAAADDFFFAAFFDDSSWW",
		//	"6A4D5F2S2W",
		//},

		//the ans can either be 62W5B or 5B62W,
		//{
		//	"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW",
		//	"62W5B",
		//},

		//ABC is just one arrangement, but we cant assert for it alone
		//{
		//	"ABC",
		//	"ABC",
		//},
	//}

	//for _, v := range sample {
	//
	//	result := example_3.Solution1(v.str)
	//
	//	assert.EqualValues(t, v.result, result)
	//
	//}

	fmt.Println(example_3.Solution1("AAAADDFFFAAFFDDSSWW"))
}