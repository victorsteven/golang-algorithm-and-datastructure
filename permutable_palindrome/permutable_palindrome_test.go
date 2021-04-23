package permutable_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutablePalindrome(t *testing.T) {
	if !PermutablePalindrome("aad") {
		t.Errorf("Want %v, got %v", true, false)
	}

	assert.True(t, PermutablePalindrome2("aad"), true)
	assert.True(t, PermutablePalindrome2("racecar"), true)
	assert.True(t, PermutablePalindrome2("rcaerca"), true)
	assert.True(t, PermutablePalindrome2("racecap"), false)
	assert.True(t, PermutablePalindrome2("rcaepca"), false)
}
