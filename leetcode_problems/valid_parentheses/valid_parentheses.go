package valid_parentheses

import "strings"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//
//
//Example 1:
//
//Input: s = "()"
//Output: true
//Example 2:
//
//Input: s = "()[]{}"
//Output: true
//Example 3:
//
//Input: s = "(]"
//Output: false
//Example 4:
//
//Input: s = "([)]"
//Output: false
//Example 5:
//
//Input: s = "{[]}"
//Output: true
//
//
//Constraints:
//
//1 <= s.length <= 104
//s consists of parentheses only '()[]{}'.


func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}



func isValid2(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}


func isValid3(str string) bool {

	for x := 0; x <= len(str)+2; x++ {

		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)

	}

	if len(str) > 0 {
		return false
	}

	return true
}

func isValid4(str string) bool {

	for x := 0; x <= len(str)+2; x++ {

		str = strings.Replace(str, "{}", "", -1)
		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
	}

	if len(str) > 0 {
		return false
	}
	return true
}
