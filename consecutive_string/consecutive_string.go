package consecutive_string

import "strings"

//ou are given an array(list) strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.
//
//Example:
//longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
//
//n being the length of the string array, if n = 0 or k > n or k <= 0 return "".
//
//Note
//consecutive strings : follow one after another without an interruption

func LongestConsec(strarr []string, k int) string {

	n := len(strarr)
	if n == 0 || k > n || k <=0 {
		return ""
	}

	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i : i+k][:], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}

	return largest

}

func LongestConsec2(strarr []string, k int) string {
	longest := ""
	for i := 0; i < len(strarr) - k + 1; i++ {
		joined := strings.Join(strarr[i:i+k], "")
		if len(joined) > len(longest) {
			longest = joined
		}
	}
	return longest
}

