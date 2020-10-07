package closest_and_smallest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosest(t *testing.T) {

	assert.Equal(t, Closest("456899 50 11992 176 272293 163 389128 96 290193 85 52"), "[(13, 9, 85), (14, 3, 176)]")

	assert.Equal(t, Closest("239382 162 254765 182 485944 134 468751 62 49780 108 54"), "[(8, 5, 134), (8, 7, 62)]")

}
