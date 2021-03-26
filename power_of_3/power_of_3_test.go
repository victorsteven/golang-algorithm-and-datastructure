package power_of_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestPower(t *testing.T) {

	assert.Equal(t, LargestPower2(4), 1)

	assert.Equal(t, LargestPower2(3), 0)

	assert.Equal(t, LargestPower2(82), 4)

	assert.Equal(t, LargestPower2(81), 3)

	assert.Equal(t, LargestPower2(7), 1)
}
