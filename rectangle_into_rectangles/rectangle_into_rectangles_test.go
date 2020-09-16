package rectangle_into_rectangles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectIntoRectangle1(t *testing.T) {

	assert.Equal(t, RectIntoRectangle4(20, 8), []string{"(16*8)", "(8*4)", "(20*8)", "(12*8)"})

	assert.Equal(t, RectIntoRectangle4(13, 5), []string{"(10*5)", "(8*5)", "(2*1)", "(3*2)", "(5*3)", "(13*5)"})

	assert.Equal(t, RectIntoRectangle4(22, 6), []string{"(12*6)", "(18*6)", "(22*6)", "(12*6)", "(16*6)", "(10*6)", "(6*4)", "(4*2)"})

}
