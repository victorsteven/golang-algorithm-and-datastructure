package main

import (
	"fmt"
	"math"
)

func cipher(str string, direction int) string {

	shift, offset := rune(math.Abs(float64(direction))), rune(26)

	runes := []rune(str)

	for i, c := range runes {

		switch  {
		case direction > 0:
			if c >= 'a' && c <= 'z' - shift || c >= 'A' && c <= 'Z' - shift {
				c = c + shift
			} else if c > 'z' - shift && c <= 'z' || c > 'Z' - shift && c <= 'Z' {
				c = c + shift - offset
			}
		case direction < 0:
			if c >= 'a' + shift && c <= 'z'  || c >= 'A' + shift && c <= 'Z' {
				c = c - shift
			} else if c >= 'a' && c < 'a' + shift || c >= 'A' + shift && c < 'A' + shift {
				fmt.Println("the one: ", string(c))
				c = c - shift + offset
			}
		}

		runes[i] = c
	}

	return string(runes)
}

//caesar("Abcd", 2) should return "Cdef"
//caesar("message", -1) should return "ldrrzfd"
//caesar("ZZ Top", 3) should return "CC Wrs"

func main() {

	result1 := cipher("Abcd", 2)
	fmt.Println(result1)

	result2 := cipher("message", -1)
	fmt.Println(result2)
	//
	result3 := cipher("ZZ Top", 3)
	fmt.Println(result3)
}
