package longest_substring

import (
	"math"
)

// Given a string s, find the length of the longest substring without repeating characters.
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

// Complexity Analysis
// Approach 1: Brute Force
//Time complexity : O(n^3)O(n
//3
// ).
//To verify if characters within index range [i, j)[i,j) are all unique, we need to scan all of them. Thus, it costs O(j - i)O(j−i) time.
//Space complexity : O(min(n, m))O(min(n,m)). We need O(k)O(k) space for checking a substring has no duplicate characters, where kk is the size of the Set. The size of the Set is upper bounded by the size of the string nn and the size of the charset/alphabet mm.

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				ans = int(math.Max(float64(ans), float64(j-i)))
			}
		}
	}
	return ans
}

func allUnique(s string, start, end int) bool {
	hash := make(map[string]bool)
	for i := start; i < end; i++ {
		char := string(s[i])
		if _, ok := hash[char]; ok {
			return false // the variable is repeating, so is not unique
		}
		hash[char] = true
	}
	return true
}

// Approach 2: Sliding Window
// The naive approach is very straightforward. But it is too slow. So how can we optimize it?
//
//In the naive approaches, we repeatedly check a substring to see if it has duplicate character. But it is unnecessary. If a substring s_{ij}s
//ij
//​
//  from index ii to j - 1j−1 is already checked to have no duplicate characters. We only need to check if s[j]s[j] is already in the substring s_{ij}s
//ij
//​
// .
//To check if a character is already in the substring, we can scan the substring, which leads to an O(n^2)O(n
//2
// ) algorithm. But we can do better.
//
//By using HashSet as a sliding window, checking if a character in the current can be done in O(1)O(1).
//
//A sliding window is an abstract concept commonly used in array/string problems. A window is a range of elements in the array/string which usually defined by the start and end indices, i.e. [i, j)[i,j) (left-closed, right-open). A sliding window is a window "slides" its two boundaries to the certain direction. For example, if we slide [i, j)[i,j) to the right by 11 element, then it becomes [i+1, j+1)[i+1,j+1) (left-closed, right-open).
//
//Back to our problem. We use HashSet to store the characters in current window [i, j)[i,j) (j = ij=i initially). Then we slide the index jj to the right. If it is not in the HashSet, we slide jj further. Doing so until s[j] is already in the HashSet. At this point, we found the maximum size of substrings without duplicate characters start with index ii. If we do this for all ii, we get our answer.

//func LengthOfLongestSubstring2(s string) int {
//	n := len(s)
//
//	hash := make(map[string]bool)
//
//	ans := 0; i := 0; j := 0
//
//	for i < n && j < n {
//		// try to extend the range [i, j]
//		if _, ok := hash[string(s[j])]; !ok {
//			hash[string(s[j + 1])] = true
//			ans = int(math.Max(float64(ans), float64(j - i)))
//			//fmt.Println("this is the ans: ", hash)
//		} else {
//			delete(hash, string(s[i+1]))
//		}
//	}
//	return ans
//}

// Approach 3: Sliding Window Optimized
func LengthOfLongestSubstring3(s string) int {

	n := len(s)
	ans := 0

	hash := make(map[string]int)

	var i int

	for j := 0; j < n; j++ {
		if _, ok := hash[string(s[i])]; ok {
			i = int(math.Max(float64(s[j]), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j-i+1)))
		hash[string(j)] = j + 1
	}
	return ans
}
