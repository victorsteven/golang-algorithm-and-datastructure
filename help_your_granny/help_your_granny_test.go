package help_your_granny_test

import (
	"efficient-algorithms/help_your_granny"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTour(t *testing.T) {
	assert.Equal(t, help_your_granny.Tour(), "")
}
