package temp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{2,5}, prime_factors(10))
}
