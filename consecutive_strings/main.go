package main

func main() {

}

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	str := ""
	if n == 0 || k > n || k <= 0 {
		return str
	}
	var max = 0
	var longest_word = ""
	for i := 0; i < n; i++ {
		if len(strarr[i]) > max {
			max = len(strarr[i])
			longest_word = strarr[i]
			if i == len(strarr)-1 {
				return longest_word
			}
		}
		for j := i + 1;  j <= k; j++ {
			str = longest_word + strarr[j]
		}
	}
	return str
}
