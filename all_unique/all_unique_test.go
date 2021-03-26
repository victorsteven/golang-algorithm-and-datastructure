package all_unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasUniqueChar(t *testing.T) {

	assert.Equal(t, false, HasUniqueChar("abasddwq"))

	assert.Equal(t, true, HashUniqueChar2("abcdef"))

	assert.Equal(t, false, HashUniqueChar3("wweede"))

}
