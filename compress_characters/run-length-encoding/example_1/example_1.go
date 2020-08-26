package example_1

import (
	"fmt"
	"strings"
)

//Given a string containing uppercase characters (A-Z), compress repeated 'runs' of the same character by storing the length of that run, and provide a function to reverse the compression. The output can be anything, as long as you can recreate the input with it.

//Input: WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW

//Output: 12W1B12W3B24W1B14W


//This deifference between this and example 2 is that for single values, add the 1 in from of the character



func Solution1(str string) string {

	count := 0
	result := ""

	for i := 0; i < len(str); i++ {

		//get the first character:
		a := string(str[i])
		//make the count to be one:
		count = 1

		//now check how many times each character occur:
		for i + 1 < len(str) && string(str[i]) == string(str[i + 1]) {
			count++
			i++
		}

		//the output:
		result += fmt.Sprintf("%d%s", count, a)

	}

	return result
}

func Solution2(str string) string {

	prev := str[0:1]
	count := 1

	result := ""

	for _, v := range strings.Split(str[1:], "") {
		if prev  == v {
			count++
		} else {
			result += fmt.Sprintf("%d%s", count, prev)
			prev = v
			count = 1 //this is very important so the previous count of another element is not added to another
		}
	}
	result += fmt.Sprintf("%d%s", count, prev)

	return result

}
