package mexican_wave

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWave(t *testing.T) {


	assert.Equal(t, Wave("hello"), []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"})

	assert.Equal(t, Wave(" x yz"), []string{" X yz", " x Yz", " x yZ"})

	assert.Equal(t, Wave(""), []string{})

	assert.Equal(t, Wave("z"), []string{"Z"})

}
