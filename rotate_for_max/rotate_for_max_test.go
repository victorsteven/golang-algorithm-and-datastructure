package rotate_for_max_test

import (
	"efficient-algorithms/rotate_for_max"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxRot(t *testing.T) {

	assert.Equal(t, rotate_for_max.MaxRot(56789), int64(68957))

}
