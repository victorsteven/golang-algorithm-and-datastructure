package cyclicRotation_test

import (
	"efficient-algorithms/codility-go/arrays/cyclicRotation"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyclicRotation(t *testing.T) {

	A := []int{3, 8, 9, 7, 6}
	k := 3

	ans := cyclicRotation.CyclicRotation(A, k)

	fmt.Println(ans)

	assert.EqualValues(t, len(A), 5)

}
