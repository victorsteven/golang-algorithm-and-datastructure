package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the lcs: ", LCS("AGGTAB", "GXTXAYB"))

	fmt.Println("this is one string: ", runner("asdf"))
}

func runner(s string) string {

	return string(s[0])
}

func LCS(s1, s2 string) string  {
	m := len(s1)
	n := len(s2)
	if m == 0 || n == 0 {
		return ""
	} else if s1[m-1] == s2[n-1] {
		return LCS(s1[:m-1], s2[:n-1]) + string(s1[m-1])
	}
	v1 := LCS(s1, s2[:n-1])
	v2 := LCS(s1[:m-1], s2)
	if len(v1) > len(v2) {
		return v1
	}
	return v2
}