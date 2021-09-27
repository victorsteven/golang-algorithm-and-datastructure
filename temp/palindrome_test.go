package temp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	assert.False(t, Palindrome("hello"))
	assert.True(t, Palindrome("ada"))
	assert.True(t, Palindrome("hannah"))

	assert.False(t, Pali("hello"))
	assert.True(t, Pali("ada"))
	assert.True(t, Pali("hannah"))
}
