package palindromestrings_test

import (
	"efficient-algorithms/palindrome/palindromestrings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	assert.True(t, palindromestrings.IsPalindrome("hannah"))
	assert.True(t, palindromestrings.IsPalindrome("hanah"))
	assert.False(t, palindromestrings.IsPalindrome("normal_http_call"))
	assert.True(t, palindromestrings.IsPalindrome("level"))

}
