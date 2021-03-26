package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords("this is the mainstream grand"))
	fmt.Println(SpinWords2("this is the mainstream grand"))
}

func SpinWords(str string) string {
	ans := ""
	for _, v := range strings.Split(str, " ") {
		if len(v) >= 5 {
			var c = strings.Split(v, "")
			for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
				c[i], c[j] = c[j], c[i]
			}
			ans += strings.Join(c, "")
			ans += " "
		} else {
			ans += v + " "
		}
	}
	return strings.TrimSpace(ans)
}

func SpinWords2(str string) string {
	ans := ""
	for _, v := range strings.Split(str, " ") {
		if len(v) >= 5 {
			result := ""
			for _, w := range v {
				result = string(w) + result
			}
			ans += result
			ans += " "
		} else {
			ans += v + " "
		}
	}
	return strings.TrimSpace(ans)
}

// this is the quickest way to reverse a word
//func Reserve(s string) (result string) {
//	for _, v := range s {
//		result = string(v) + result
//	}
//	return result
//}
