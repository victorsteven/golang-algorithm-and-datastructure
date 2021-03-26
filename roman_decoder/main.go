package main

import (
	"fmt"
)

func main() {
	fmt.Println(Decode("MMVIII"))
	fmt.Println(Decode2("MMVIII"))
}

var ROMAN = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Decode(roman string) int {
	result := 0
	last_digit := 1000
	for _, v := range roman {
		fmt.Printf("The type %T\n", v) //this is a rune === int32
		digit := ROMAN[v]
		if last_digit < digit {
			result -= 2 * last_digit
		}
		last_digit = digit
		result += digit
	}
	return result
}

func Decode2(roman string) int {
	var intern []int
	var result int
	for _, r := range roman {
		intern = append(intern, ROMAN[r])
	}
	for i := 1; i < len(intern); i++ {
		if intern[i-1] >= intern[i] {
			result += intern[i-1]
		} else {
			result -= intern[i-1]
		}
	}
	result += intern[len(intern)-1]

	return result
}
