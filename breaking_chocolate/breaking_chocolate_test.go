package breaking_chocolate_test

import (
	"efficient-algorithms/breaking_chocolate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakChocolate(t *testing.T) {

	assert.Equal(t, breaking_chocolate.BreakChocolate(2, 1), 1)
	assert.Equal(t, breaking_chocolate.BreakChocolate(3, 1), 2)
}
