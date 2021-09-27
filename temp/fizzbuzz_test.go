package temp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "12fizz4buzzfizz78fizz", fizzbuzz(10))
	assert.Equal(t, "12fizz4buzzfizz78fizzbuzz11fizz1314", fizzbuzz(15))
}