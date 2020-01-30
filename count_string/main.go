package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countString("aaabbbvraqq"))
}

func countString(s string) []string {
	number_values := []string{}
	for _, v := range s {
		ans := strings.Count(strings.ToLower(s), string(v))
		ansMod := fmt.Sprintf("%v-%v", string(v), ans)
		number_values = append(number_values, ansMod)
	}
	newArray := []string{}
	keys := make(map[string]bool)
	for _, v := range number_values {
		if value := keys[v]; !value {
			keys[v] = true
			newArray = append(newArray, v)
		}
	}
	return newArray
}
