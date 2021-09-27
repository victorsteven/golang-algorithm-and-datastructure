package temp

func prime_factors(n int) []int {
	divisor := 2
	var factors []int
	for n > 1 {
		if n % divisor == 0 {
			factors = append(factors, divisor)
			n = n/divisor
		} else {
			divisor++
		}
	}
	return factors
}
