package reversesentences_test

import (
	"efficient-algorithms/reverses/reversesentences"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSentence(t *testing.T) {
	assert.Equal(t, reversesentences.ReverseSentence("hello"), "hello")
	assert.Equal(t, reversesentences.ReverseSentence("this is a sentence"), "sentence a is this")
	assert.Equal(t, reversesentences.ReverseSentence("      hello      "), "      hello      ")
	assert.Equal(t, reversesentences.ReverseSentence("   this is a sentence"), "sentence a is this   ")
	assert.Equal(t, reversesentences.ReverseSentence("What is happening?"), "happening? is What")
	assert.Equal(t, reversesentences.ReverseSentence("Hello! How are you?"), "you? are How Hello!")
}
