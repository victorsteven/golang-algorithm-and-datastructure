package upside_down_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpsideDown(t *testing.T) {

	assert.Equal(t, UpsideDown("0", "10"), uint64(3))
	assert.Equal(t, UpsideDown("0", "10"), uint64(3))
	assert.Equal(t, UpsideDown("6", "25"), uint64(2))
	assert.Equal(t, UpsideDown("10", "100"), uint64(4))
	assert.Equal(t, UpsideDown("100", "1000"), uint64(12))
	//assert.Equal(t, UpsideDown("100000","12345678900000000"), uint64(718650))
}
