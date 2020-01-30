package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//fmt.Println("this iz a sample ans 2: ", matchGet("helloMike"))
	//fmt.Println("this iz a sample ans: ", matching("Hello"))
	//fmt.Println("the uppers: ", "This is the sample: ", removeDelimeterAndLowerCase("This is for the CAPS"))
	//fmt.Println("the uppers: ", "This is the sample: ", removeDelimeterAndUpperCase("This is for the CAPS"))

	//fmt.Println(mix("First string", "Second"))
	s1 := "my&friend&Paul has heavy hats! &"
	s2 := "my friend John has many many friends &"
	fmt.Println(mix(s1, s2))

}

func countLetter(str, letter string) int {
	count := 0
	for _, v := range str {
		if string(v) == letter {
			count++
		}
	}
	return count
}

func mix(s1, s2 string) string {
	var allLowerCase = regexp.MustCompile(`[^[:lower:]]`)
	s1  = allLowerCase.ReplaceAllString(s1, "")
	s2  = allLowerCase.ReplaceAllString(s2, "")

	output := []string{}

	for _, v := range s1 {
		subArr := []string{}
		if strings.Contains(s2, string(v)) {
			countLet_s1 := countLetter(s1, string(v))
			countLet_s2 := countLetter(s2, string(v))
			if countLet_s1 < 2 && countLet_s2 < 2 {
				continue
			} else {
				if countLet_s1 > countLet_s2 {
					subArr = append(subArr, "1:")
					subArr = append(subArr, strings.Repeat(string(v), countLet_s1))
				} else if countLet_s2 > countLet_s1 {
					subArr = append(subArr, "2:")
					subArr = append(subArr, strings.Repeat(string(v), countLet_s2))
				} else {
					subArr = append(subArr, "=:")
					subArr = append(subArr, strings.Repeat(string(v), countLet_s1))
				}
			}
		}
		output = append(output, subArr...)
	}
	var newArray []string
	keys := make(map[string]bool)

	for i, v := range output {
		if  v != "2:" && v != "1:" && v != "=:" {
			if _, value := keys[v]; !value {
				keys[v] = true
				newArray = append(newArray, output[i-1]) //add the 2:, 1: or =: before the characters
				newArray = append(newArray, v)
				newArray = append(newArray, " ")
			}
		}
	}
	//convert the slice of string to string
	str := strings.Join(newArray, "")
	newest := strings.Replace(str, " ", "/", -1)
	//remove the last item in the string:
	newest = newest[:len(newest)-1]

	return newest
}


//return a bool
//func matching(str string) bool {
//	var isLetter = regexp.MustCompile(`^[a-z]+$`).MatchString
//
//	return isLetter(str)
//}

func removeDelimeterAndLowerCase(str string) string {
	var allUpperCase = regexp.MustCompile(`[^[:upper:]]`)
	return allUpperCase.ReplaceAllString(str, "")
}

func removeDelimeterAndUpperCase(str string) string {
	var allLowerCase = regexp.MustCompile(`[^[:lower:]]`)
	return allLowerCase.ReplaceAllString(str, "")
}

