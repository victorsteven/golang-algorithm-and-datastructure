package all_unique

import "strings"

// Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.
//
//The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g. 'a' and 'A' are considered different characters.

func HasUniqueChar(s string) bool {

	hash := make(map[string]int)
	for _, v := range s {
		hash[string(v)]++
	}

	for _, v := range hash {
		if v > 1 {
			return false
		}
	}
	return true
}

func HashUniqueChar2(str string) bool {
	seen := make(map[rune]bool)
	for _, c := range str {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func HashUniqueChar3(str string) bool {
	for _, letter := range str {
		if strings.Count(str, string(letter)) > 1 {
			return false
		}
	}
	return true
}
