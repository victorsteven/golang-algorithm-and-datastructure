package main

import "fmt"

//Generate a string consisting of characters ‘a’ and ‘b’ that satisfy the given conditions

//Given two integers A and B, the task is to generate and print a string str such that:
//
//str must only contain the characters ‘a’ and ‘b’.
//str has length A + B and the occurrence of character ‘a’ is equal to A and the occurrence of character ‘b’ is equal to B
//The sub-strings “aaa” or “bbb” must not occur in str.
//Note that for the given values of A and B, a valid string can always be generated.

//Examples:
//
//Input: A = 1, B = 2
//Output: abb
//“abb”, “bab” and “bba” are all valid strings.
//
//Input: A = 4, B = 1
//Output: aabaa

func generateString(A int, B int) string {

	rt := ""

	for A > 0 || B > 0 {
		//more 'b' , append "bba"
		if A < B {
			B--
			if 0 < B {
				rt += "b"
			}
			A--
			if 0 < A {
				rt += "a"
			}
			//More 'a', append "aab"
		} else if B < A {
			A--
			if 0 < A {
				rt += "a"
			}
			B--
			if 0 < B {
				rt += "b"
			}
			//equal number of 'a' and 'b', append "ab"
		} else {
			A--
			if 0 < A {
				rt += "a"
			}
			B--
			if 0 < B {
				rt += "b"
			}
		}
	}
	return rt
}

func main() {
	ans := generateString(2, 6)
	fmt.Println(ans)
}
