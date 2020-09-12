package encrypt_this

import (
	"fmt"
	"strings"
)

//Encrypt this!
//
//You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:
//
//Your message is a string containing space separated words.
//You need to encrypt each word in the message using the following rules:
//The first letter needs to be converted to its ASCII code.
//The second letter needs to be switched with the last letter
//Keepin' it simple: There are no special characters in input.
//Examples:
//EncryptThis("Hello") == "72olle"
//EncryptThis("good") == "103doo"
//EncryptThis("hello world") == "104olle 119drlo"



func EncryptThis(text string) string {

	if len(text) == 0 {
		return text
	}

	str := strings.Split(text, " ")

	strEncrypt := ""

	for _, v := range str {

		strEncrypt += fmt.Sprintf("%s ", encryptThis(v))

	}

	return strings.TrimSpace(strEncrypt)
}

func encryptThis(text string) string {

	r := []rune(text)

	if len(text)-1 > 0 {
		r[1], r[len(text)-1] = r[len(text)-1], r[1]
	}

	rString := string(r)

	output := fmt.Sprintf("%d%s", r[0], rString[1:len(rString)])

	return output
}



//func EncryptThis(text string) string {
//
//	res := ""
//
//	for _, word := range strings.Split(text, " ") {
//		for i, letter := range word {
//			switch i {
//			case 0: res += strconv.Itoa(int(letter))
//			case 1: res += string(word[len(word)-1])
//			case len(word)-1: res += string(word[1])
//			default:
//				res += string(letter)
//			}
//		}
//		res += " "
//	}
//	return strings.TrimSpace(res)
//}

//func EncryptThis(text string) string {
//	if len(text) == 0 {return text}
//	words := strings.Split(text, " ")
//	for i, s := range words {
//		r := []rune(s)
//		n := len(r)-1
//		if n > 1 { r[1], r[n] = r[n], r[1] }
//		r = append([]rune(strconv.FormatInt(int64(r[0]), 10)), r[1:]...)
//		words[i] = string(r)
//	}
//	return strings.Join(words, " ")
//}

