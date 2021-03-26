package numbers_in_order_test

import (
	"efficient-algorithms/numbers_in_order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInAscOrder(t *testing.T) {

	assert.True(t, numbers_in_order.InAscOrder([]int{1, 2, 4, 7}))
	assert.True(t, numbers_in_order.InAscOrder([]int{1, 2, 3, 4, 5}))
	assert.False(t, numbers_in_order.InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))

}
