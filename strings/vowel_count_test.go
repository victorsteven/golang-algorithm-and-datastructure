package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVowelCount(t *testing.T) {
	assert.Equal(t, 6, vowel_count("aweruhedbage"))
}
