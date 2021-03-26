package twice_linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDblLinear(t *testing.T) {

	assert.Equal(t, DblLinear(50), 175)
	assert.Equal(t, DblLinear(100), 447)
	assert.Equal(t, DblLinear2(500), 3355)
	assert.Equal(t, DblLinear2(1000), 8488)

}
