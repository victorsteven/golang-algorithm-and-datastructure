package baseE91

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {

}


func TestEncode(t *testing.T) {

	assert.Equal(t, Encode([]byte("test")), []byte("fPNKd"))
}