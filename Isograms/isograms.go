package Isograms

import "strings"

//An isogram is a word that has no repeating letters, consecutive or non-consecutive. Implement a function that determines whether a string that contains only letters is an isogram. Assume the empty string is an isogram. Ignore letter case.
//
//isIsogram "Dermatoglyphics" == true
//isIsogram "aba" == false
//isIsogram "moOse" == false -- ignore letter case

func Isogram(str string) bool {

	if len(str) == 0 {
		return true
	}

	hashMap := make(map[rune]bool)

	for _, v := range strings.ToLower(str) {
		if ok := hashMap[v]; ok {
			return false
		}
		//insert
		hashMap[v] = true
	}

	return true
}
