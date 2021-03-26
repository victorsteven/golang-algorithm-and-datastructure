package alphabet_symmetry

import (
	"strings"
	"unicode"
)

//Consider the word "abode". We can see that the letter a is in position 1 and b is in position 2. In the alphabet, a and b are also in positions 1 and 2. Notice also that d and e in abode occupy the positions they would occupy in the alphabet, which are positions 4 and 5.
//
//Given an array of words, return an array of the number of letters that occupy their positions in the alphabet for each word. For example,
//
//solve(["abode","ABc","xyzD"]) = [4, 3, 1]
//See test cases for more examples.
//
//Input will consist of alphabet characters, both uppercase and lowercase. No spaces.
//
//Good luck!
//
//If you like this Kata, please try:

func Solve(slice []string) (result []int) {

	for _, str := range slice {
		count := 0
		for i, ch := range strings.ToLower(str) {
			if int(ch-'a') == i {
				count++
			}
		}
		result = append(result, count)
	}
	return
}

func Solve2(slice []string) []int {

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	results := make([]int, len(slice))
	//loop string
	for index, str := range slice {
		//loop character
		for charIndex, char := range str {
			if unicode.ToLower(char) == rune(alphabet[charIndex]) {
				results[index]++
			}
		}
	}
	return results
}
