package longest_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	//assert.Equal(t, 3, LengthOfLongestSubstring("abcabcbb"))
	//assert.Equal(t, 1, LengthOfLongestSubstring("bbbbb"))

	//assert.Equal(t, 3, LengthOfLongestSubstring2("abcabcbb"))
	assert.Equal(t, 1, LengthOfLongestSubstring3("bbbbb"))
	assert.Equal(t, 3, LengthOfLongestSubstring3("abcabcbb"))
}
