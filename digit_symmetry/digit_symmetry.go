package digit_symmetry

import (
	"strconv"
)

//Consider the number 1176 and its square (1176 * 1176) = 1382976. Notice that:
//
//the first two digits of 1176 form a prime.
//the first two digits of the square 1382976 also form a prime.
//the last two digits of 1176 and 1382976 are the same.
//Given two numbers representing a range (a, b), how many numbers satisfy this property within that range? (a <= n < b)


//Example
//solve(2, 1200) = 1, because only 1176 satisfies this property within the range 2 <= n < 1200. See test cases for more examples. The upper bound for the range will not exceed 1,000,000.
//
//Good luck!
//
//If you like this Kata, please try:

func DigitSymmetry(a, b int) int {
	n := 0
	for i := a; i < b; i++ {
		x := strconv.Itoa(i)
		y := strconv.Itoa(i * i)

		if len(x) > 2 && len(y) > 2 {
			if x[len(x)-2:] != y[len(y)-2:] {
				continue
			}
			if !isPrime(x[:2]) {
				continue
			}
			if !isPrime(y[:2]) {
				continue
			}
			n++
		}
	}
	return n
}

func isPrime(x string) bool {

	n, _ := strconv.Atoi(x)

	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true

}
