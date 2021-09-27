package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseArray(t *testing.T) {
	arr := []int{1,2,3,4,5,6}
	result := reverse_array(arr)
	assert.Equal(t, result, []int{6,5,4,3,2,1})
}