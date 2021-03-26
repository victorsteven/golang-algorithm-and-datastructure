package upside_down_2

import (
	"strconv"
)

//Welcome to the Challenge Edition of Upside-Down Numbers
//In the original kata by @KenKamau you were limited to integers below 2^17. Here, you will be given strings of digits of up to 42 characters in length (upper bound is 10^42 - 1).
//
//Your task is essentially the same, but an additional challenge is creating a fast, efficient solution.
//
//Input:
//Your function will receive two strings, each comprised of digits representing a positive integer. These two values will represent the upper and lower bounds of a range.
//
//Output:
//Your function must return the number of valid upside down numbers within the range of the two input arguments, including both upper and lower bounds.
//
//What is an Upside-Down Number?
//An upside down number is an integer that appears the same when rotated 180 degrees, as illustrated below.
//
//1961 - OK,
//88 - OK,
//101 - OK,
//25 - No
//
//Example:
//
//var x,y string = "0","25"
//upsideDown(x,y); //4
////the valid numbers in the range are 0, 1, 8, and 11

func UpsideDown(n1, n2 string) uint64 {
	// your code goes here. you can do it!

	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)

	var count uint64

	outcome := make(chan bool)

	for i := num1; i < num2; i++ {
		go isUpSideDown(i, outcome)
	}

	for i := num1; i < num2; i++ {
		if <-outcome == true {
			count++
		}
	}

	return count
}

func isUpSideDown(x int, outcome chan bool) {

	//convert the number to a string:
	str := strconv.Itoa(x)

	//reverse the number:
	strRev := reverse(str)

	result := make([]rune, 0)

	for _, v := range strRev {
		result = append(result, exchange(v))
	}

	value, _ := strconv.Atoi(string(result))

	outcome <- value == x
}

func reverse(str string) string {

	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func exchange(r rune) rune {

	switch r {
	case '0':
		return '0'
	case '1':
		return '1'
	case '6':
		return '9'
	case '8':
		return '8'
	case '9':
		return '6'
	}
	return ' '
}
