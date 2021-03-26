package example_2

import (
	"fmt"
	"strings"
)

//Run-length encoding (RLE) is a simple "compression algorithm" (an algorithm which takes a block of data and reduces its size, producing a block that contains the same information in less space). It works by replacing repetitive sequences of identical data items with short "tokens" that represent entire sequences. Applying RLE to a string involves finding sequences in the string where the same character repeats. Each such sequence should be replaced by a "token" consisting of:
//
//the number of characters in the sequence the repeating character If a character does not repeat, it should be left alone.
//
//For example, consider the following string:

//qwwwwwwwwweeeeerrtyyyyyqqqqwEErTTT

//After applying the RLE algorithm, this string is converted into:
//
//q9w5e2rt5y4qw2Er3T

func Solution1(str string) string {

	count := 0
	result := ""

	for i := 0; i < len(str); i++ {
		var a = string(str[i])
		count = 1

		for i+1 < len(str) && string(str[i]) == string(str[i+1]) {
			count++
			i++
		}

		if count == 1 {
			result += fmt.Sprintf("%s", a)
		} else {
			result += fmt.Sprintf("%d%s", count, a)
		}
	}

	return result
}

func Solution2(str string) string {

	prev := str[0:1]
	count := 1

	result := ""

	for _, v := range strings.Split(str[1:], "") {
		if prev == v {
			count++
		} else {
			if count == 1 {
				result += fmt.Sprintf("%s", prev)
			} else {
				result += fmt.Sprintf("%d%s", count, prev)
			}
			prev = v
			count = 1
		}
	}

	if count == 1 {
		result += fmt.Sprintf("%s", prev)
	} else {
		result += fmt.Sprintf("%d%s", count, prev)
	}

	return result

}
