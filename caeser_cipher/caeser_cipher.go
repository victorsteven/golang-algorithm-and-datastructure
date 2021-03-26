package caeser_cipher

import "math"

// Caesar Cipher
// (description more or less taken from Wikipedia)
//
//	In cryptography, a Caesar cipher, also known as Caesar's cipher,
//	the shift cipher, Caesar's code or Caesar shift, is one of the
//	simplest and most widely known encryption techniques. It is a type
//	of substitution cipher in which each letter in the plaintext is
//	replaced by a letter some fixed number of positions down the
//	alphabet. For example, with a left shift of 3, D would be replaced
//	by A, E would become B, and so on. The method is named after Julius
//	Caesar, who used it in his private correspondence.

// cipher takes in the text to be ciphered along with the direction that
// is being taken; -1 means encoding, +1 means decoding.

//func Cipher(text string, direction int) string {
//
//	//shift is the number of letters to move to right or left
//	//offset is the size of the alphabet in this case the plain ASCII
//
//	shift, offset := rune(3), rune(26)
//
//	runes  := []rune(text)
//
//	for i, char := range runes {
//		//if the letter is not in the range [1..25], the offset defined above is added or subsctracted.
//
//		switch direction {
//		case -1: //encoding
//			if char >= 'a' + shift && char <= 'z' || char >= 'A' + shift && char <= 'Z' {
//				char = char - shift
//			} else if char >= 'a' && char < 'a' + shift || char >= 'A' && char < 'A' + shift {
//				char = char - shift + offset
//			}
//		case +1: //decoding
//			if char >= 'a' && char <= 'z' - shift || char >= 'A' && char <= 'Z' - shift {
//				char = char + shift
//			} else if char > 'z' - shift && char <= 'z' || char > 'Z' - shift && char <= 'Z' {
//				char = char + shift - offset
//			}
//		}
//
//		runes[i] = char
//	}
//	return string(runes)
//}

//
//func encode(text string) string {
//	return Cipher(text, -1)
//}
//
//func decode(text string) string {
//	return Cipher(text, +1)
//}

func Cipher(text string, direction int) string {

	//shift is the number of letters to move to right or left
	//offset is the size of the alphabet in this case the plain ASCII

	//note, the shift should only be positive numbers
	shift, offset := rune(math.Abs(float64(direction))), rune(26)

	runes := []rune(text)

	for i, char := range runes {
		//if the letter is not in the range [1..25], the offset defined above is added or subsctracted.

		switch {

		case direction < 0:
			if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case direction > 0:
			if char >= 'a' && char <= 'z'-shift || char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' || char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		runes[i] = char
	}
	return string(runes)
}
