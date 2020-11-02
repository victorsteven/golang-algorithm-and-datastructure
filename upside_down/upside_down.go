package upside_down

import (
	"strconv"
	"strings"
)

//Consider the numbers 6969 and 9116. When you rotate them 180 degrees (upside down), these numbers remain the same. To clarify, if we write them down on a paper and turn the paper upside down, the numbers will be the same. Try it and see! Some numbers such as 2 or 5 don't yield numbers when rotated.
//
//Given a range, return the count of upside down numbers within that range. For example, solve(0,10) = 3, because there are only 3 upside down numbers >= 0 and < 10. They are 0, 1, 8.
//
//More examples in the test cases.
//
//Good luck!
//
//If you like this Kata, please try
//
//Simple Prime Streaming
//
//Life without primes
//
//Please also try the performance version of this kata at Upside down numbers - Challenge Edition



func Solution(a, b int) int {

	count := 0

	for i := a; i < b; i++ {
		if upsideDown(i) {
			count++
		}
	}
	return count
}

func upsideDown(x int) bool {

	str := strconv.Itoa(x)

		strRev := reverse(str)
		strArr := strings.Split(strRev, "")

		result := ""
		for _, v := range strArr {
			result += upsideNumbers(v)
		}

		value, _ := strconv.Atoi(result)

		return value == x
}

func upsideNumbers(s string) string {
	switch s {
	case "0":
		return "0"
	case "1":
		return "1"
	case "6":
		return "9"
	case "8":
		return "8"
	case "9":
		return "6"
	}
	return ""
}

func reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}


