package valid_parentheses_test

import (
	"efficient-algorithms/valid_parentheses"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {

	assert.False(t, valid_parentheses.ValidParentheses(")"))
	assert.False(t, valid_parentheses.ValidParentheses2(")(()))"))
	assert.True(t, valid_parentheses.ValidParentheses2("()"))
	assert.True(t, valid_parentheses.ValidParentheses3("(())((()())())"))

	assert.True(t, valid_parentheses.ValidParentheses4("(())((()())())"))

}
