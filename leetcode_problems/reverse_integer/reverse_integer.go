package reverse_integer

import (
	"math"
	"strconv"
)

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
//Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
//Example 1:
//
//Input: x = 123
//Output: 321
//Example 2:
//
//Input: x = -123
//Output: -321
//Example 3:
//
//Input: x = 120
//Output: 21
//Example 4:
//
//Input: x = 0
//Output: 0
//
//
//Constraints:
//
//-2^31 <= x <= 2^31 - 1

func reverse(x int) int {

	runes := make([]rune, 0)
	str := strconv.Itoa(x)
	if x < 0 {
		str = str[1:] // remove the minus sign
		runes = []rune(str)
	} else {
		runes = []rune(str)
	}

	for i, j := 0, len(runes)-1;  i < j; i, j = i + 1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	y, _ := strconv.Atoi(string(runes))

	if float64(y) < math.Pow(-2, 31) || float64(y) > math.Pow(2, 31) {
		return 0
	}

	if x < 0 {
		return y * -1
	}
	return y
}
