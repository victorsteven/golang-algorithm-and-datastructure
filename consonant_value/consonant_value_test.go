package consonant_value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {

	assert.Equal(t, Solve("zodiacs"), 26)
	assert.Equal(t, Solve("uaoczei"), 29)
	assert.Equal(t, Solve2("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"), 143)
	assert.Equal(t, Solve2("codewars"), 37)

	assert.Equal(t, Solve("this is the man at the balcony"), 39)

}
