package longest_substring_without_repeatition

import (
	"math"
)

// 3. Longest Substring Without Repeating Characters
//Medium
//
//13976
//
//723
//
//Add to List
//
//Share
//Given a string s, find the length of the longest substring without repeating characters.
//
//
//
//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//Example 4:
//
//Input: s = ""
//Output: 0
//
//
//Constraints:
//
//0 <= s.length <= 5 * 104
//s consists of English letters, digits, symbols and spaces.


// Using Brute Force:
// Complexity Analysis
//
//Time complexity : O(n^3)O(n
//3
// ).
//
//To verify if characters within index range [i, j)[i,j) are all unique, we need to scan all of them. Thus, it costs O(j - i)O(j−i) time.


func lengthOfLongestSubstring(s string) int {
   n := len(s)
   res := 0
   for i := 0; i < n; i++ {
   		for j := i; j < n; j++ {
   			if checkRepetition(s, i, j) {
   				res = int(math.Max(float64(res), float64(j-i+1)))
			}
		}
   }
   return res
}

func checkRepetition(s string, start, end int) bool {
	chars := make(map[uint8]int, 0)
	for i := start; i <= end; i++ {
		c := s[i]
		chars[c]++
		if chars[c] > 1 {
			return false
		}
	}
	return true
}


func lengthOfLongestSubstring2(s string) int {

	chars := make(map[uint8]int, 0)

	left := 0
	right := 0
	res := 0

	for right > len(s) {
		r := s[right]
		chars[r]++

		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}
		res = int(math.Max(float64(res), float64(right-left+1)))

		right++
	}
	return res
}

func lengthOfLongestSubstring3(s string) int {

	chars := make(map[uint8]int, 0)
	n := len(s)
	ans := 0

	for j, i := 0,0; j < n; j++ {
		if _, ok := chars[s[j]]; ok {
			i = int(math.Max(float64(chars[s[j]]), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j - i + 1)))
		chars[s[j]] = j + 1
	}
	return ans
}


func lengthOfLongestSubstring4(s string) int {
	chars := make(map[uint8]int, 0)
	ans := 0
	for j, i := 0,0; j < len(s); j++ {
		if _, ok := chars[s[j]]; ok {
			i = int(math.Max(float64(chars[s[j]]), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j - i + 1)))
		chars[s[j]] = j + 1
	}
	return ans
}


