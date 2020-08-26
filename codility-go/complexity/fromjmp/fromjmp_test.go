package fromjmp_test

import (
	"efficient-algorithms/codility-go/complexity/fromjmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromJmp(t *testing.T) {

	ans1 := fromjmp.FromJmp(10, 85, 30)
	ans2 := fromjmp.FromJmp(1, 5, 2)

	assert.EqualValues(t, ans1, 3)
	assert.EqualValues(t, ans2, 2)
}
