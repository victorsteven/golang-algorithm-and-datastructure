package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("this is the ans: ", FindShort("normal_http_call my name isdfs assdf"))
}

func FindShort(s string) int {
	if len(s) <= 0 {
		return 0
	}
	words := strings.Split(s, " ")
	lowestWord := words[0]
	for _, v := range words {
		if len(v) < len(lowestWord) {
			lowestWord = v
		}
	}
	return len(lowestWord)
}
