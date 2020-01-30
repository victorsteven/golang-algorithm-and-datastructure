package main

import (
	"fmt"
	"strconv"
)

//Write a program that prints the integers from   1   to   100   (inclusive).
//
//But:
//
//for multiples of three,   print   Fizz     (instead of the number)
//for multiples of five,   print   Buzz     (instead of the number)
//for multiples of both three and five,   print   FizzBuzz     (instead of the number)
//
//The   FizzBuzz   problem was presented as the lowest level of comprehension required to illustrate adequacy.

func main() {
fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {

	result := []string{}
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}