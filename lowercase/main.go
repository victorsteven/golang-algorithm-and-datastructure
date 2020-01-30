package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(alpha("THEgenaway mine own HELLObo"))
}

//get only the lowercase letters
func alpha(str string) string {
	//re := regexp.MustCompile("[]")
	var allLowerCase = regexp.MustCompile(`[^[:lower:]]`)

	str = allLowerCase.ReplaceAllString(str, " ")

	return str
}