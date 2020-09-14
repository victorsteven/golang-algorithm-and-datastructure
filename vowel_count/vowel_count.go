package vowel_count

import (
	"regexp"
	"strings"
)

//Return the number (count) of vowels in the given string.
//
//We will consider a, e, i, o, u as vowels for this Kata (but not y).
//
//The input string will only consist of lower case letters and/or spaces.

func GetCount(str string) (count int) {

	count = 0
	for _, v := range strings.ToLower(str) {
		switch v {
		case 'a':
			count++
		case 'e':
			count++
		case 'i':
			count++
		case 'o':
			count++
		case 'u':
			count++
		}
	}

	return count
}


func GetCount2(str string) (count int) {
	for _, c := range str  {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func GetCount3(strn string) int {
	count := 0
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)
	}
	return count
}

func GetCount4(str string) int {

	r := regexp.MustCompile("[aeiou]")

	vowels := r.FindAllString(str, -1)

	return len(vowels)
}