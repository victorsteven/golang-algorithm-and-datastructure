package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	arr := []int{1,2,2,4,5,5,2,6,6,5}
	expected := []int{1,2,4,5,6}
	result := remove_duplicate(arr)

	assert.Equal(t, expected, result)
}
