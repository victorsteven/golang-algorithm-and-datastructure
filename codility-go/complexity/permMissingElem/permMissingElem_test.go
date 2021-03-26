package permMissingElem_test

import (
	"efficient-algorithms/codility-go/complexity/permMissingElem"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerMissingElem(t *testing.T) {

	A := []int{2, 3, 1, 5}
	ans := permMissingElem.PerMissingElem(A)
	assert.EqualValues(t, ans, 4)

	A2 := []int{1, 2, 3, 4, 5, 7}
	ans2 := permMissingElem.PerMissingElem(A2)
	assert.EqualValues(t, ans2, 6)
}
