package valid_parentheses

import (
	"regexp"
	"strings"
)

//Write a function called that takes a string of parentheses, and determines if the order of the parentheses is valid. The function should return true if the string is valid, and false if it's invalid.

//Examples
//"()"              =>  true
//")(()))"          =>  false
//"("               =>  false
//"(())((()())())"  =>  true

//Constraints
//0 <= input.length <= 100

func ValidParentheses(parens string) bool {

	if len(parens) >= 100 {
		return false
	}

	b := 0
	for _, c := range parens {
		if c == '(' {
			b++
		} else {
			b--
		}
		if b < 0 {
			return false
		}
	}
	return b == 0
}


func ValidParentheses2(parens string) bool {

	if len(parens) >= 100 {
		return false
	}

	_, err := regexp.Compile(parens)

	return err == nil
}


func ValidParentheses3(parens string) bool {

	if len(parens) >= 100 {
		return false
	}

	for strings.Contains(parens, "()") {
		parens = strings.Replace(parens, "()", "", -1)
	}
	return len(parens) == 0
}

func ValidParentheses4(parens string) bool {

	n := 0

	for _, char := range parens {
		switch char {
		case '(': n++
		case ')': n--
		}
		if n < 0 {
			return false
		}
	}
	return n == 0
}