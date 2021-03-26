package upside_down

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	assert.Equal(t, Solution(0, 10), 3)
	assert.Equal(t, Solution(10, 100), 4)
	assert.Equal(t, Solution(1000, 10000), 20)
	assert.Equal(t, Solution(10000, 15000), 6)
	assert.Equal(t, Solution(15000, 20000), 9)
	assert.Equal(t, Solution(60000, 70000), 15)
	assert.Equal(t, Solution(60000, 130000), 55)

}

//func TestReverseNow(t *testing.T) {
//
//	assert.Equal(t, ReverseNow("green"), "neerg")
//}
