package two_to_one

import (
	"fmt"
	"sort"
	"strings"
)

//Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters,
//
//each taken only once - coming from s1 or s2.
//Examples:
//a = "xyaabbbccccdefww"
//b = "xxxxyyyyabklmopq"
//longest(a, b) -> "abcdefklmopqwxy"
//
//a = "abcdefghijklmnopqrstuvwxyz"
//longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"

func distinct(s string) string {

	dis := make(map[rune]bool)

	str := ""

	for _, v := range s {
		if val := dis[v]; !val {
			dis[v] = true
			str += fmt.Sprintf("%c", v)
		}
	}
	return str
}

func TwoToOne(s1 string, s2 string) string {

	totalString := s1 + s2

	finalString := distinct(totalString)

	strArr := strings.Split(finalString, "")

	sort.Strings(strArr)

	return strings.Join(strArr, "")

}

func TwoToOne2(s1, s2 string) string {

	strArr := strings.Split(s1+s2, "")
	sort.Strings(strArr)

	result := ""

	for _, s := range strArr {
		if !strings.Contains(result, s) {
			result += s
		}
	}
	return result
}

func TwoToOne3(s1, s2 string) string {

	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
			result += char
		}
	}
	return result
}
