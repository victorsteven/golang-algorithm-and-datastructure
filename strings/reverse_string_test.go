package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "yebba", reverse_string("abbey"))
}
