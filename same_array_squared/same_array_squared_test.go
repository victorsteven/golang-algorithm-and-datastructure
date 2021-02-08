package same_array_squared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComp(t *testing.T) {

	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	assert.True(t, Comp(a, b))

	a2 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b2 := []int{132, 14641, 20736, 361, 25921, 361, 20736, 361}
	assert.False(t, Comp(a2, b2))

	var a3 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var b3 = []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	assert.True(t, Comp(a3, b3))

	a4 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b4 := []int{11*21, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	assert.False(t, Comp(a4, b4))

	b5 := []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	assert.False(t, Comp(nil, b5))
}
