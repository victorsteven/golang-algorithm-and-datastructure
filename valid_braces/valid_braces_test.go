package valid_braces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidBraces(t *testing.T) {

	assert.Equal(t, ValidBraces("(){}[]"), true)

	assert.Equal(t, ValidBraces("([{}])"), true)

	assert.Equal(t, ValidBraces("(}"), false)

	assert.Equal(t, ValidBraces("[(])"), false)
	//
	//assert.Equal(t, ValidBraces("[({})](]"), false)



}


