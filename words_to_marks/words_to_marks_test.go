package words_to_marks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordsToMarks(t *testing.T) {

	assert.Equal(t, WordsToMarks("attitude"), 100)
	assert.Equal(t, WordsToMarks2("friends"), 75)
	assert.Equal(t, WordsToMarks3("family"), 66)
}
