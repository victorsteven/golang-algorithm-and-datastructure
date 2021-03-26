package fix_string_case

import (
	"regexp"
	"strings"
	"unicode"
)

//In this Kata, you will be given a string that may have mixed uppercase and lowercase letters and your task is to convert that string to either lowercase only or uppercase only based on:
//
//make as few changes as possible.
//if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
//For example:
//
//solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
//solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
//solve("coDE") = "code". Upper == lowercase. Change all to lowercase.
//More examples in test cases. Good luck!

func Solve(str string) string {

	upperReg := regexp.MustCompile(`[A-Z][^A-Z]*`)

	up := upperReg.FindAllString(str, -1)

	lowerReg := regexp.MustCompile(`[a-z][^a-z]*`)

	low := lowerReg.FindAllString(str, -1)

	if len(up) > len(low) {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}
}

func Solve2(str string) string {

	uppers := 0

	for _, rune := range str {
		if unicode.IsUpper(rune) {
			uppers++
		}
	}
	if uppers > len(str)/2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
