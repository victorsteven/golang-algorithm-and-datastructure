package growing_plant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrowingPlant(t *testing.T) {

	assert.Equal(t, GrowingPlant(100, 10, 910), 10)
	assert.Equal(t, GrowingPlant(10, 9, 4), 1)
	assert.Equal(t, GrowingPlant2(5, 2, 6), 2)
	assert.Equal(t, GrowingPlant3(72, 2, 369), 6)

}
