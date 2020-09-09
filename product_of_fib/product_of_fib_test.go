package product_of_fib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductFib(t *testing.T) {

	assert.Equal(t, ProductFib(714), [3]uint64{21, 34, 1})
	assert.Equal(t, ProductFib(800), [3]uint64{34, 55, 0})

}
