package main

import (
	"fmt"
	"regexp"
)

//In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.
//
//The string has the following conditions to be alphanumeric:
//
//At least one character ("" is not valid)
//Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
//No whitespaces / underscore

func main() {
	fmt.Println(alphanumeric("THEgenaw@#$@*(ay mine own HELLObo!@#$#"))
	fmt.Println(alphanumeric("THEsuf dsufniudf jdsfniusdf"))
	fmt.Println(alphanumeric("THEsuf df"))

}

func alphanumeric(str string) bool {
	// ^ means negation
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return r.MatchString(str)
}

