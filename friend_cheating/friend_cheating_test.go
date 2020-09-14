package friend_cheating_test

import (
	"efficient-algorithms/friend_cheating"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemovNb(t *testing.T) {

	ans := friend_cheating.RemovNb(26)
	fmt.Println("Our ans: ", ans)

	assert.Equal(t, friend_cheating.RemovNb(101), [][2]uint64{{55, 91}, {91, 55}})
	assert.Equal(t, friend_cheating.RemovNb(102), [][2]uint64{{70, 73}, {73, 70}})
	assert.Equal(t, friend_cheating.RemovNb(110), [][2]uint64{{70, 85}, {85, 70}})

}
