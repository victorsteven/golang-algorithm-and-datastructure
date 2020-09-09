package digit_symmetry

import (
	"strconv"
)

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
