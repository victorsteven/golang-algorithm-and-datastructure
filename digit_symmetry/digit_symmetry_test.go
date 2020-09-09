package digit_symmetry_test

import (
	"efficient-algorithms/digit_symmetry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigitSymmetry(t *testing.T) {

	assert.Equal(t, digit_symmetry.DigitSymmetry(2, 1200), 1)

	assert.Equal(t, digit_symmetry.DigitSymmetry(2,100000), 247)

	assert.Equal(t, digit_symmetry.DigitSymmetry(2,1000000), 2549)

	assert.Equal(t, digit_symmetry.DigitSymmetry(100000,1000000), 2302)

}
