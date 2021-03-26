package passphrases

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Everyone knows passphrases. One can choose passphrases from poems, songs, movies names and so on but frequently they can be guessed due to common cultural references. You can get your passphrases stronger by different means. One is the following:
//
//choose a text in capital letters including or not digits and non alphabetic characters,
//
//shift each letter by a given number but the transformed letter must be a letter (circular shift),
//replace each digit by its complement to 9,
//keep such as non alphabetic and non digit characters,
//downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
//reverse the whole result.
//#Example:
//
//your text: "BORN IN 2015!", shift 1
//
//1 + 2 + 3 -> "CPSO JO 7984!"
//
//4 "CpSo jO 7984!"
//
//5 "!4897 Oj oSpC"
//
//With longer passphrases it's better to have a small and easy program. Would you write it?

func PlayPass(s string, n int) string {

	// the complement of digits table;
	digitHash := map[int]int{
		0: 9, 1: 8, 2: 7, 3: 6, 4: 5, 5: 4, 6: 3, 7: 2, 8: 1, 9: 0,
	}

	var result string

	// first transform:
	for _, v := range s {

		isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(string(v))

		digit, err := strconv.Atoi(string(v))
		if err == nil {
			//is a digit, so apply the rule:
			result += strconv.Itoa(digitHash[digit])
		} else if isLetter {
			theRune := v + int32(n)
			var newletter string
			if theRune > 90 {
				newletter = string('a' + int32(n))
			} else {
				newletter = string(v + int32(n))
			}
			//newletter := string(v + int32(n))
			fmt.Println("the current no: ", v+int32(n), newletter)
			if len(result)%2 == 0 {
				result += strings.ToUpper(newletter)
			} else if len(result)%2 != 0 {
				result += strings.ToLower(newletter)
			}
		} else {
			result += string(v)
		}
	}

	// MY GRANMA CAME FROM NY ON THE 23RD OF APRIL 2015

	newResult := Reverse(result)

	return newResult
}

func Reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func transform(r rune, shift int32) rune {
	if unicode.IsLower(r) {
		r = 'a' + (r-'a'+shift)%26
	}
	if unicode.IsUpper(r) {
		r = 'A' + (r-'A'+shift)%26
	}
	if unicode.IsDigit(r) {
		r = '9' - r + '0'
	}
	return r
}

func PlayPass2(s string, n int) string {
	p := ""
	for i, r := range s {
		r = transform(r, int32(n))
		if i%2 == 0 && unicode.IsLetter(r) {
			r = unicode.ToUpper(r)
		} else {
			r = unicode.ToLower(r)
		}
		p += string(r)
	}
	return p
}
