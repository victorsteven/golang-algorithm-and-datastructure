package split_string_distinct

//Given a string s. Let k be the maximum number of partitions possible of the given string with each partition starts with distinct character. The task is to find the number of ways string s can be split into k partition (non-empty) such that each partition start with distinct character.

func Solution(s string) int {

	count := make([]int, 26) //the length of the string is the length of alphabets

	//check how many times the same character occurs:
	for _, v := range s {
		count[v-'a']++
	}

	//making frequency of first character of string equal to 1
	count[s[0]-'a'] = 1

	// Finding the product of frequency of occurrence of each character.
	ans := 1

	//get the others:
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			ans *= count[i]
		}
	}

	return ans
}
