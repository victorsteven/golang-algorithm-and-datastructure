package decimal_to_factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDec2FactString(t *testing.T) {

	assert.Equal(t, Dec2FactString(36288000), "A0000000000")
	assert.Equal(t, Dec2FactString(2982), "4041000")

	assert.Equal(t, FacsString2Dec("341010"), 463)
	assert.Equal(t, FacsString2Dec("4042100"), 2990)

}
