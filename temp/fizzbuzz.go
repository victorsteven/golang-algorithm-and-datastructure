package temp

import (
	"strconv"
)

func fizzbuzz(n int) string {
	result := ""
	for i := 1; i < n; i++ {
		switch  {
		case i % 5 == 0 && i % 3 == 0:
			result += "fizzbuzz"
		case i % 3 == 0:
			result += "fizz"
		case i % 5 == 0:
			result += "buzz"
		default:
			result += strconv.Itoa(i)
		}
	}
	return result
}
