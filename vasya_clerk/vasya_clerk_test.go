package vasya_clerk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTickets(t *testing.T) {

	assert.Equal(t, Tickets([]int{25, 25, 50}), "YES")
	assert.Equal(t, Tickets2([]int{25, 100}), "NO")
	assert.Equal(t, Tickets3([]int{25, 25, 50, 50, 100}), "NO")

}
