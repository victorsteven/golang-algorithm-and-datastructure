package deadfish_swim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert.Equal(t, []int{8, 64}, Parse("iiisdoso"))
	assert.Equal(t, []int{0, 1}, Parse("idoiido"))
	assert.Equal(t, []int{1,4,25}, Parse("isoisoiso"))
}
