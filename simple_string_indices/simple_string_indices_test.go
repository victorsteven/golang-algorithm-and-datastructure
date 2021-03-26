package simple_string_indices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	result1, _ := Solution("((1)23(45))(aB)", 0)
	result2, _ := Solution2("((1)23(45))(aB)", 11)

	assert.Equal(t, int(result1), 10)
	assert.Equal(t, int(result2), 14)

}
