package main

import "math"

func main() {

}

//You are given an array strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.
//Example:
//
//longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
//
//n being the length of the string array, if n = 0 or k > n or k <= 0 return "".

func MxDifLg(a1 []string, a2 []string) int {

	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var max = 0
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			var diff = int(math.Abs(float64(len(a1[i])) - float64(len(a2[j]))))
			if diff > max {
				max = diff
			}
		}
	}
	return max
}
