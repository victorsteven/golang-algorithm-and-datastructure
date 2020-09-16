package variadic_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	assert.Equal(t, Sum(3,4,5), 12)
	assert.Equal(t, Sum(1,4,5), 10)
	assert.Equal(t, Sum(5), 5)

	nums := []int{1,2,3,4,5}

	assert.Equal(t, Sum(nums...), 15)

}
