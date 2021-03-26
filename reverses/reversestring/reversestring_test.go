package reversestring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, ReverseString("normal_http_call"), "olleh")
}
