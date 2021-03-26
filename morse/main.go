package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(decodeMorse(".... . -.--   .--- ..- -.. ."))
	fmt.Println(decodeMorse(".... . -.--   .--- ..- -.. ."))
}

var MORSE_CODE = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

//func decodeMorse(morseCode string) string {
//	decoded := ""
//	morseWords := strings.Split(strings.TrimSpace(morseCode), "   ")
//	for _, word := range morseWords {
//		morseWord := strings.Split(word, " ")
//		for _, char := range morseWord {
//			decoded += MORSE_CODE[char]
//		}
//		decoded += " "
//	}
//	return strings.TrimSpace(decoded)
//}

func decodeMorse(morseCode string) string {
	decoded := ""
	morseWords := strings.Split(morseCode, "   ")
	for _, v := range morseWords {
		morseChars := strings.Split(v, " ")
		for _, char := range morseChars {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}
