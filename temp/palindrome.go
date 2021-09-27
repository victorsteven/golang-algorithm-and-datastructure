package temp

const (
	capitalA rune = 'A'
	capitalZ rune = 'Z'
	smallA rune = 'a'
	smallZ rune = 'z'
)

func Palindrome(s string) bool {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1

	for i < j {
		if !isCharacter(runes[i]) {
			i++
			continue
		}
		if !isCharacter(runes[j]) {
			j--
			continue
		}
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isCharacter(char rune) bool {
	if char >= capitalA && char <= capitalZ {
		return true
	}
	if char >= smallA && char <= smallZ {
		return true
	}
	return false
}

func isCaseInsensitive(char rune) bool {
	if char >= smallA && char <= capitalZ {
		return true
	}
	return false
}

func Pali(s string) bool {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1

	for i < j {
		if isCaseInsensitive(runes[i]) {
			i++
			continue
		}
		if isCaseInsensitive(runes[j]) {
			j--
			continue
		}
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}


