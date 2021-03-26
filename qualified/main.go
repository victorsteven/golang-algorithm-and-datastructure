package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
//Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.
//
//Specification
//solution(number)
//Finds the sum of multiples of 3 or 5 that is less than the provided number
//
//Parameters
//number: int - Maximum number to check against
//
//Return Value
//int - Sum of all multiples of either 3 or 5
//
//Examples
//number	Return Value
//10	23
//200	9168

func Solution(n int) int {

	result := 0

	for i := 1; i <= n; i++ {
		if i < n {
			switch {
			case i%3 == 0 && i%5 == 0:
				result += i
			case i%3 == 0:
				result += i
			case i%5 == 0:
				result += i
			}
		}
	}
	return result
}

func Solution2(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if i < n {
			switch {
			case i%3 == 0 && i%5 == 0:
				result += i
			case i%3 == 0:
				result += i
			case i%5 == 0:
				result += i
			}
		}
	}
	return result
}

func Solution3(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if i < n {
			switch {
			case i%3 == 0 && i%5 == 0:
				result += i
			case i%3 == 0:
				result += i
			case i%5 == 0:
				result += i
			}
		}
	}
	return result
}

//You are provided a string containing a list of positive integers separated by a space (" "). Take each value and calculate the sum of its digits, which we call it's "weight". Then return the list in ascending order by weight, as a string joined by a space.
//
//For example 99 will have "weight" 18, 100 will have "weight"
//1 so in the ouput 100 will come before 99.
//
//Specification
//orderWeight(strng)
//Parameters
//strng: String - String of digits to be summed and put in order
//
//Return Value
//String - A string of digits ordered by their "weight"
//
//Example:
//
//"56 65 74 100 99 68 86 180 90" ordered by numbers weights becomes:
//"100 180 90 56 65 74 68 86 99"
//
//When two numbers have the same "weight", let's consider them to be strings and not numbers:
//
//100 is before 180 because its "weight" (1) is less than the one of 180 (9)
//and 180 is before 90 since, having the same "weight" (9) it comes before as a string.
//
//All numbers in the list are positive integers and the list can be empty.

func OrderByWeight(strng string) string {
	items := strings.Split(strng, " ")
	sort.SliceStable(items, func(i, j int) bool {
		x := weight(items[i])
		y := weight(items[j])

		if (x - y) == 0 {
			return items[i] < items[j]
		}

		return weight(items[i]) <= weight(items[j])
	})

	return strings.Join(items, " ")
}

func OrderByWeight2(str string) string {
	items := strings.Split(str, " ")
	sort.SliceStable(items, func(i, j int) bool {
		x := weight2(items[i])
		y := weight2(items[j])

		if (x - y) == 0 {
			return items[i] < items[j]
		}
		return weight2(items[i]) <= weight(items[j])
	})
	return strings.Join(items, " ")
}

func OrderByWeight3(str string) string {
	items := strings.Split(str, " ")
	sort.SliceStable(items, func(i, j int) bool {
		x := weight3(items[i])
		y := weight3(items[j])

		if (x - y) == 0 {
			return items[i] < items[j]
		}
		return weight3(items[i]) <= weight3(items[j])

	})

	return strings.Join(items, " ")
}

func weight3(str string) int {
	result := 0

	for _, n := range str {
		intV, _ := strconv.Atoi(string(n))
		result += intV
	}
	return result
}
func weight2(str string) int {
	result := 0
	for _, n := range str {
		intV, _ := strconv.Atoi(string(n))
		result += intV
	}
	return result
}

func weight(str string) int {

	result := 0
	for _, v := range str {
		intV, _ := strconv.Atoi(string(v))
		result += intV
	}

	return result
}

// Write a function called validBraces that takes a string of braces, and determines if the order of the braces is valid. validBraces should return true if the string is valid, and false if it's invalid.
//
//All input strings will be nonempty, and will only consist of open parentheses '(' , closed parentheses ')', open brackets '[', closed brackets ']', open curly braces '{' and closed curly braces '}'.
//
//What is considered Valid?
//
//A string of braces is considered valid if all braces are matched with the correct brace. For example:
//
//'(){}[]' and '([{}])' would be considered valid, while '(}', '[(])', and '[({})](]' would be considered invalid.
//
//Specification
//validBraces(braces)
//Checks if the brace order is valid
//
//Parameters
//braces: String - A string representation of an order of braces
//
//Return Value
//Boolean - Returns true if order of braces are valid
//
//Examples:
//Input	Output
//validBraces( "(){}[]" )	true
//validBraces( "(}" )	false
//validBraces( "[(])" )	false
//validBraces( "([{}])" )	true

func ValidBraces(str string) bool {

	for x := 0; x <= len(str)+2; x++ {

		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)

	}

	if len(str) > 0 {
		return false
	}

	return true
}

func ValidBraces2(str string) bool {
	for x := 0; x <= len(str)+2; x++ {
		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)
	}
	if len(str) > 0 {
		return false
	}
	return true
}

func ValidBraces3(str string) bool {
	str = strings.Replace(str, "()", "", -1)
	str = strings.Replace(str, "[]", "", -1)
	str = strings.Replace(str, "{}", "", -1)

	if len(str) > 0 {
		return false
	}
	return true
}

func OrderByWeight4(str string) string {
	items := strings.Split(str, " ")
	sort.SliceStable(items, func(i, j int) bool {
		x := weight4(items[i])
		y := weight4(items[j])

		if (x - y) == 0 {
			return items[i] < items[j]
		}

		return weight4(items[i]) <= weight4(items[j])
	})

	return strings.Join(items, " ")
}

func weight4(str string) int {
	result := 0

	for _, n := range str {
		intV, _ := strconv.Atoi(string(n))
		result += intV
	}
	return result
}

func Solution4(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if i < n {
			switch {
			case i%3 == 0 && i%5 == 0:
				result += i
			case i%3 == 0:
				result += i
			case i%5 == 0:
				result += i
			}
		}
	}
	return result
}

func main() {

	fmt.Println(Solution4(10))
	fmt.Println("valid brace: ", ValidBraces3("(){}[]"))
	fmt.Println("final: ", OrderByWeight4("56 65 74 100 99 68 86 180 90"))
}
