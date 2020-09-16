package two_to_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoToOne(t *testing.T) {

	assert.Equal(t, TwoToOne("aretheyhere", "yestheyarehere"), "aehrsty")

	assert.Equal(t, TwoToOne2("loopingisfunbutdangerous", "lessdangerousthancoding"), "abcdefghilnoprstu")

	assert.Equal(t, TwoToOne3("loopingisfunbutdangerous", "lessdangerousthancoding"), "abcdefghilnoprstu")

}
