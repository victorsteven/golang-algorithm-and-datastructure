package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(print_error("aaabbbbhaijjjm"))
	//fmt.Println(print_error("aaaxbbbbyyhwawiwjjjwwm"))

	fmt.Println(print_error2("aaabbbbhaijjjm"))
	fmt.Println(print_error2("aaaxbbbbyyhwawiwjjjwwm"))
}

func print_error(str string) string  {
	if len(str) <= 0 {
		return ""
	}
	match := 0
	standard := "abcdefghijklm"
	//standard := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m"}
	for _, v := range str {
		for _, w := range standard {
			if v == w {
				match++
			}
		}
	}
	num := 0
	strLen := len(str)
	if  strLen > match {
		num = strLen - match
	} else {
		num = match - len(str)
	}
	return strconv.Itoa(num) + "/" + strconv.Itoa(strLen)
}

//Solution 2
func print_error2(s string) string {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'm' {
			continue
		} else {
			count++
		}
	}
	return fmt.Sprintf("%d/%d", count, len(s))
}
