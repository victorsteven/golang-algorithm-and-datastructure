package consecutive_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestLongestConsec(t *testing.T) {

	assert.Equal(t, LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2), "abigailtheta")
	assert.Equal(t, LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1), "oocccffuucccjjjkkkjyyyeehh")
	assert.Equal(t, LongestConsec([]string{"itvayloxrp","wkppqsztdkmvcuwvereiupccauycnjutlv","vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2), "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	assert.Equal(t, LongestConsec([]string{}, 3), "")

}
