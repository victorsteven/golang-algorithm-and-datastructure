package main

import (
	"fmt"
	"strings"
)

//Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character. Example 1: Input: S = "ab#c", T = "ad#c" Output: true Explanation: Both S and T become "ac". Example 2: Input: S = "ab##", T = "c#d#" Output: true Explanation: Both S and T become "". Example 3: Input: S = "a##c", T = "…

func main() {

	result := solution("#####ab", "ab")

	fmt.Println(result)
}


func filterString(s string) string {

	str := []string{}

	for i, v := range s {
		if v == '#' {
			if len(str) > 0 {
				str = str[:i]
			}
		} else {
			str  = append(str, string(v))
		}
	}

	return strings.Join(str, "")
}


func solution(S, T string) bool {
	if filterString(S) != filterString(T) {
		return false
	}
	return true
}
