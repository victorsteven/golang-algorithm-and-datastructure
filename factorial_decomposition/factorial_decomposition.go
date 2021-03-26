package factorial_decomposition

import (
	"fmt"
	"strings"
)

//The aim of the kata is to decompose n! (factorial n) into its prime factors.
//
//Examples:
//
//n = 12; decomp(12) -> "2^10 * 3^5 * 5^2 * 7 * 11"
//since 12! is divisible by 2 ten times, by 3 five times, by 5 two times and by 7 and 11 only once.
//
//n = 22; decomp(22) -> "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"
//
//n = 25; decomp(25) -> 2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23
//Prime numbers should be in increasing order. When the exponent of a prime is 1 don't put the exponent.
//
//Notes
//
//the function is decomp(n) and should return the decomposition of n! into its prime factors in increasing order of the primes, as a string.
//factorial can be a very big number (4000! has 12674 digits, n will go from 300 to 4000).
//In Fortran - as in any other language - the returned string is not permitted to contain any redundant trailing whitespace: you can use dynamically allocated character strings.

func Decomp(n int) string {

	//THIS WORKS FOR SMALL NUMBERS
	//nFac :=  getFactorialInt(n)
	//
	//fmt.Println("these are the primes: ", nFac)
	//
	//primeHash := make(map[int]int, 0)
	//var divisor = 2
	//for nFac  > 1 {
	//	if nFac % divisor == 0 {
	//		primeHash[divisor]++
	//		nFac = nFac / divisor
	//	} else {
	//		divisor++
	//		if divisor == 3 {
	//			fmt.Println("The current number", nFac)
	//		}
	//	}
	//}
	//
	////sort the keys:
	//keys := make([]int, 0, len(primeHash))
	//for k := range primeHash {
	//	keys = append(keys, k)
	//}
	//sort.Ints(keys)

	//result := []string{}
	//
	//for _, k := range keys {
	//	if primeHash[k] != 1 {
	//		result = append(result, fmt.Sprintf("%d^%d", k, primeHash[k]))
	//	} else {
	//		result = append(result, fmt.Sprintf("%d", k))
	//	}
	//}
	//return strings.Join(result, " * ")

	//FOR BIG NUMBERS
	primeList := []int{}
	primesTable := map[int]int{}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeList = append(primeList, i)
			primesTable[i] = 1
		} else {
			var temp = i
			for j := range primesTable {
				for temp%j == 0 {
					primesTable[j]++
					temp /= j
				}
			}
		}
	}

	result := []string{}

	for _, k := range primeList {
		if primesTable[k] != 1 {
			result = append(result, fmt.Sprintf("%d^%d", k, primesTable[k]))
		} else {
			result = append(result, fmt.Sprintf("%d", k))
		}
	}
	return strings.Join(result, " * ")
}

//for big numbers
func getFactorial(n float64) float64 {

	if n < 2 {
		return 1
	}

	return n * getFactorial(n-1)
}

func getFactorialInt(n int) int {

	if n < 2 {
		return 1
	}

	return n * getFactorialInt(n-1)
}

//func isPrime(i int) bool {
//	ans := big.NewInt(int64(i)).ProbablyPrime(0)
//	return ans
//}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
