package gap_in_primes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func  TestGap(t *testing.T) {

	//assert.Equal(t, AllPossiblePrimes(20), []int{2, 3, 5, 7, 11, 13, 17, 19})

	//assert.Equal(t, Gap(2, 3, 50), []int{3, 5})
	//assert.Equal(t, Gap(2,100,110), []int{101, 103})
	//assert.Equal(t, Gap(4,100,110), []int{103, 107})
	assert.Equal(t, Gap(6,100,110), nil)
	//assert.Equal(t, Gap(8,300,400), []int{359, 367})
	//assert.Equal(t, Gap(10,300,400), []int{337, 347})

	//dotest(2,100,110, []int{101, 103})
	//dotest(4,100,110, []int{103, 107})
	//dotest(6,100,110, nil)
	//dotest(8,300,400, []int{359, 367})
	//dotest(10,300,400, []int{337, 347})

}
