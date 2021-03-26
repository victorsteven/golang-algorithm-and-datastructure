package pile_of_cube

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNb(t *testing.T) {
	assert.Equal(t, FindNb(1071225), 45)
}
