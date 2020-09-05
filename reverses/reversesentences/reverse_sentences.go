package reversesentences

import "efficient-algorithms/reverses/reversestring"

func ReverseSentence(sentence string) string {

	runes := []rune(sentence)

	size := len(runes)

	//reverse entire string
	reversestring.Reverse(runes, 0, size-1)

	start := 0
	end := 0

	for start < size && end < size {
		if runes[start] == ' ' {
			start++
			end++
			continue
		}

		if runes[end] != ' ' {
			end++
			continue
		}

		reversestring.Reverse(runes, start, end-1)

		start = end
	}

	reversestring.Reverse(runes, start, end-1)

	return string(runes)

}