package count_divisors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivisors(t *testing.T) {

	assert.Equal(t, 3, Divisors(4))
	assert.Equal(t, 2, Divisors(5))
	assert.Equal(t, 6, Divisors(12))
	assert.Equal(t, 8, Divisors(30))
}