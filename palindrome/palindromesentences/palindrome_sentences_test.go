package palindromesentences_test

import (
	"efficient-algorithms/palindrome/palindromesentences"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	assert.True(t, palindromesentences.IsPalindrome("hannah"))
	assert.True(t, palindromesentences.IsPalindrome("too hot to hoot."))
	assert.False(t, palindromesentences.IsPalindrome("normal_http_call"))

	assert.True(t, palindromesentences.IsPalindromeCaseInsensitive("Do geese see God?"))
	assert.True(t, palindromesentences.IsPalindromeCaseInsensitive("Was it a car or a cat I saw?"))
}
