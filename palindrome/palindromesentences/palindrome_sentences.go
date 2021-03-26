package palindromesentences

import "strings"

//IsPalindrome determines if the input is a palindrome

const (
	capitalA = 'A'
	capitalZ = 'Z'
	smallA   = 'a'
	smallZ   = 'z'
)

func IsPalindromeCaseInsensitive(s string) bool {

	//change to lower case
	sentence := strings.ToLower(s)

	runes := []rune(sentence)
	length := len(runes)

	i := 0
	j := length - 1

	for i < j {
		if !isCharacterCaseInsensitive(runes[i]) {
			i++
			continue
		}
		if !isCharacterCaseInsensitive(runes[j]) {
			j--
			continue
		}
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func IsPalindrome(sentence string) bool {

	runes := []rune(sentence)

	length := len(runes)

	i := 0
	j := length - 1

	for i < j {
		if !isCharacter(runes[i]) {
			i++
			continue
		}

		if !isCharacter(runes[j]) {
			j--
			continue
		}

		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isCharacter(character rune) bool {
	if character >= capitalA && character <= capitalZ {
		return true
	}
	if character >= smallA && character <= smallZ {
		return true
	}
	return false
}

func isCharacterCaseInsensitive(character rune) bool {

	if character >= smallA && character <= smallZ {
		return true
	}
	return false
}
