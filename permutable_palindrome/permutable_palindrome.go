package permutable_palindrome

func PermutablePalindrome(s string) bool {

	hashMap := make(map[rune]int, 0)

	for _, v := range s {
		hashMap[v]++
	}

	odd := 0
	for _, v := range hashMap {
		if v % 2 != 0 {
			odd++
		}
	}
	if odd > 1 {
		return false
	}
	return true
}

func PermutablePalindrome2(s string) bool {

	hashMap := make(map[rune]int, 0)
	odd := 0
	for _, v := range s {
		if _, ok := hashMap[v]; ok {
			if hashMap[v] > 1 && hashMap[v] % 2 != 0 {
				odd++
			}
		}
		hashMap[v]++
	}
	if odd > 1 {
		return false
	}
	return true
}
