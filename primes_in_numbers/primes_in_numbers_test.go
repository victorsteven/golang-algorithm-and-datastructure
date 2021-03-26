package primes_in_numbers_test

import (
	"efficient-algorithms/primes_in_numbers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimeFactors(t *testing.T) {

	assert.Equal(t, "(2**5)(5)(7**2)(11)", primes_in_numbers.PrimeFactors(86240))

	assert.Equal(t, "(2**5)(5)(7**2)(11)", primes_in_numbers.PrimeFactors(86240))

	ans := primes_in_numbers.PrimeFactors(614031)

	fmt.Println("ans : ", ans)

}
