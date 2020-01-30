package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(string_ending("abc", "bc"))
	fmt.Println(string_ending("abc", "d"))

	fmt.Println(string_ending2("abc", "bc"))

}

func string_ending(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
func string_ending2(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	fmt.Println("this is the print first: ", str[len(ending):])
	fmt.Println("this is the print: ", str[len(str)-len(ending):])
	return str[len(str)-len(ending):] == ending
}
