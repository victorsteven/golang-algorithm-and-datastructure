package simple_string_indices

import (
	"errors"
)

//In this Kata, you will be given a string with brackets and an index of an opening bracket and your task will be to return the index of the matching closing bracket. Both the input and returned index are 0-based except in Fortran where it is 1-based. An opening brace will always have a closing brace. Return -1 if there is no answer (in Haskell, return Nothing; in Fortran, return 0; in Go, return an error)
//
//Examples
//solve("((1)23(45))(aB)", 0) = 10 // the opening brace at index 0 matches the closing brace at index 10
//solve("((1)23(45))(aB)", 1) = 3
//solve("((1)23(45))(aB)", 2) = -1 // there is no opening bracket at index 2, so return -1
//solve("((1)23(45))(aB)", 6) = 9
//solve("((1)23(45))(aB)", 11) = 14
//solve("((>)|?(*'))(yZ)", 11) = 14
//Input will consist of letters, numbers and special characters, but no spaces. The only brackets will be ( and ).
//
//More examples in the test cases.
//
//Good luck!

func Solution(str string, idx uint) (uint, error) {

	if str[idx] != '(' {
		return 0, errors.New("Not an opening bracket")
	}

	stack := 1
	for i := idx + 1; i < uint(len(str)); i++ {
		if str[i] == '(' {
			stack++
		} else if str[i] == ')' {
			stack--
		}
		if stack == 0 {
			return i, nil
		}
	}

	return 0, nil
}

func Solution2(s string, i uint) (uint, error) {
	if s[i] == '(' {
		open := 0
		for ; int(i) < len(s); i++ {
			switch s[i] {
			case '(': open++
			case ')': open--
			}
			if open == 0 {
				return i, nil
			}
		}
	}
	return 0, errors.New("Not an opening bracket")
}
