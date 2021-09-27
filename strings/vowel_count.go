package strings

func vowel_count(s string) int {
	sum := 0
	for _, v := range s {
		switch v {
		case 'a','e','i','o','u':
			sum++
		}
	}
	return sum
}
