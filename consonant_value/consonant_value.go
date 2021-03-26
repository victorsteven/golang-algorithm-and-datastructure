package consonant_value

import (
	"regexp"
)

//Given a lowercase string that has alphabetic characters only and no spaces, return the highest value of consonant substrings. Consonants are any letters of the alphabet except "aeiou".
//
//We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.
//
//For example, for the word "zodiacs", let's cross out the vowels. We get: "z o d ia cs"
//
//-- The consonant substrings are: "z", "d" and "cs" and the values are z = 26, d = 4 and cs = 3 + 19 = 22. The highest is 26.
//solve("zodiacs") = 26
//
//For the word "strength", solve("strength") = 57
//-- The consonant substrings are: "str" and "ngth" with values "str" = 19 + 20 + 18 = 57 and "ngth" = 14 + 7 + 20 + 8 = 49. The highest is 57.
//For C: do not mutate input.
//
//More examples in test cases. Good luck!

//func Solve(str string) int {
//
//	reg := regexp.MustCompile(`[aeiou]`)
//
//	consonantStr := reg.ReplaceAllString(str, "-")
//
//	consoArr := strings.Split(consonantStr, "-")
//
//	numArr := []int{}
//
//	for _, s := range consoArr {
//		num := 0
//		for _, ch := range s {
//			num += int(ch) - 'a' + 1
//		}
//		numArr = append(numArr, num)
//	}
//
//	max := 0
//
//	for _, v := range numArr {
//		if max < v {
//			max = v
//		}
//	}
//
//	return max
//
//}

func Solve(str string) int {

	reg := regexp.MustCompile(`[aeiou]`)

	consoArr := reg.Split(str, -1)

	numArr := []int{}

	for _, s := range consoArr {
		num := 0
		for _, ch := range s {
			num += int(ch) - 'a' + 1
		}
		numArr = append(numArr, num)
	}

	max := 0

	for _, v := range numArr {
		if max < v {
			max = v
		}
	}

	return max

}

func Solve2(str string) (maximum int) {

	var sum int

	for _, r := range str {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			sum = 0
			continue
		}
		sum += int(r) - 'a' + 1

		if sum > maximum {
			maximum = sum
		}
	}
	return maximum
}
