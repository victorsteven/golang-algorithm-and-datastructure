package palindrome_number

import (
	"fmt"
	"math"
	"strconv"
)

// Given an integer x, return true if x is palindrome integer.
//
//An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
//
//
//
//Example 1:
//
//Input: x = 121
//Output: true
//Example 2:
//
//Input: x = -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//
//Input: x = 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//Example 4:
//
//Input: x = -101
//Output: false
//
//
//Constraints:
//
//-231 <= x <= 231 - 1
//
//
//Follow up: Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {

	s := strconv.Itoa(x)
	runes := make([]rune, 0)

	if x < 0 {
		s = s[1:]
		runes = []rune(s)
	} else {
		runes = []rune(s)
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	result := string(runes)

	resultNum, _ := strconv.Atoi(result)
	if x < 0 {
		resultNum = resultNum * -1
	}
	if float64(resultNum) < math.Pow(-2, 31) || float64(resultNum) > math.Pow(2, 31) {
		return false
	}

	if x < 0 {
		result = fmt.Sprintf("%v%v", result, "-")
	}

	return result == s
}
