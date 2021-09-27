package temp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 13, fibonacci(7))
	assert.Equal(t, 13, loopFib(7))
}
