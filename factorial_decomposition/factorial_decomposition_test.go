package factorial_decomposition_test

import (
	"efficient-algorithms/factorial_decomposition"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecomp(t *testing.T) {

	//assert.Equal(t, "2^15 * 3^6 * 5^3 * 7^2 * 11 * 13 * 17",  factorial_decomposition.Decomp(17))
	//assert.Equal(t, "2^3 * 3 * 5",  factorial_decomposition.Decomp(5))

	//assert.Equal(t, "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19",  factorial_decomposition.Decomp(22))

	assert.Equal(t, "2^11 * 3^5 * 5^2 * 7^2 * 11 * 13",  factorial_decomposition.Decomp(14))

	//assert.Equal(t, "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23",  factorial_decomposition.Decomp(22))
	//assert.Equal(t, "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23",  factorial_decomposition.Decomp1(25))

	//assert.Equal(t, "2^15 * 3^6 * 5^3 * 7^2 * 11 * 13 * 17",  factorial_decomposition.Decomp1(17))

}
