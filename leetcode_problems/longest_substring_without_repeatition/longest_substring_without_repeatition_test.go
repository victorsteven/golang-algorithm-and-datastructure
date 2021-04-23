package longest_substring_without_repeatition

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {

	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Errorf("Want %d, got %d", 3, lengthOfLongestSubstring4("abcabcbb"))
	}
}
