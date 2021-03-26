package word_values

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameValue(t *testing.T) {

	//input := []string{"abc","abc","abc","abc"}
	//result := NameValue(input)
	//assert.Equal(t, result, []int{6,12,18,24})

	input2 := []string{"abc", "abc abc"}

	result2 := NameValue(input2)

	assert.Equal(t, result2, []int{6, 24})
}
