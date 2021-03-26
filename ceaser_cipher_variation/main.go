package main

import (
	"fmt"
	"strings"
)

//In this country soldiers are poor but they need a certain level of secrecy for their communications so, though they do not know Caesar cypher, they reinvent it in the following way.
//
//They use ASCII, without really knowing it, but code only letters a-z and A-Z. Other characters are kept such as.
//
//They change the "rotate" each new message. This "rotate" is a prefix for their message once the message is coded. The prefix is built of 2 letters, the second one being shifted from the first one by the "rotate", the first one is the first letter, after being downcased, of the uncoded message.
//
//For example if the "rotate" is 2, if the first letter of the uncoded message is 'J' the prefix should be 'jl'.
//
//To lessen risk they cut the coded message and the prefix in five pieces since they have only five runners and each runner has only one piece.
//
//If possible the message will be evenly split between the five runners; if not possible, parts 1, 2, 3, 4 will be longer and part 5 shorter. The fifth part can have length equal to the other ones or shorter. If there are many options of how to split, choose the option where the fifth part has the longest length, provided that the previous conditions are fulfilled. If the last part is the empty string don't put this empty string in the resulting array.
//
//For example, if the coded message has a length of 17 the five parts will have lengths of 4, 4, 4, 4, 1. The parts 1, 2, 3, 4 are evenly split and the last part of length 1 is shorter. If the length is 16 the parts will be of lengths 4, 4, 4, 4, 0. Parts 1, 2, 3, 4 are evenly split and the fifth runner will stay at home since his part is the empty string and is not kept.
//
//Could you ease them in programming their coding?
//
//Example with shift = 1 :
//
//message : "I should have known that you would have a perfect answer for me!!!"
//
//code : => ["ijJ tipvme ibw", "f lopxo uibu z", "pv xpvme ibwf ", "b qfsgfdu botx", "fs gps nf!!!"]
//
//By the way, maybe could you give them a hand to decode?
//
//Caesar cipher : see Wikipedia

const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lower = "abcdefghijklmnopqrstuvwxyz"
const alphaLen = len(upper)

var shiftVal int

// mapFn - used to call string.Map()
func mapFn(r rune) rune {
	if strings.ContainsRune(upper, r) || strings.ContainsRune(lower, r) {
		var baseChar = int(upper[0])
		var charSet = upper
		if strings.ContainsRune(lower, r) {
			baseChar = int(lower[0])
			charSet = lower
		}
		zeroBased := int(r) - int(baseChar) // Make zero-based ('A' or 'a' == 0)
		zeroBased += shiftVal
		if zeroBased < 0 {
			zeroBased += alphaLen
		}
		zeroBased %= alphaLen
		r = rune(charSet[zeroBased])
	}
	return r
}

// Encode - encode a message into an array of strings
func Encode(s string, shift int) []string {
	// your code
	shiftVal = shift % 26 // Set the global shift amount
	keyBase := strings.ToLower(string(s[0]))
	keyStr := keyBase + strings.Map(mapFn, string(keyBase))
	var encodedStr = keyStr + strings.Map(mapFn, s)
	var encodedStrLen = len(encodedStr)
	var sectionF, section = float64(encodedStrLen) / 5.0, encodedStrLen / 5
	if sectionF-float64(section) > 0 {
		section++
	}
	fmt.Println("String:", s)
	fmt.Println("Shift:", shiftVal)
	fmt.Println("encodedStr:", encodedStr)
	fmt.Println("encodedStrLen:", encodedStrLen)
	fmt.Println("sectionF, section, rem", sectionF, section, encodedStrLen-4*section)
	var rslt []string
	for i := 0; i < 4; i++ {
		rslt = append(rslt, encodedStr[i*section:(i+1)*section])
	}
	fmt.Println("result len:", len(rslt))
	if encodedStrLen > 4*section {
		rslt = append(rslt, encodedStr[4*section:])
	}
	fmt.Println("result len:", len(rslt))
	return rslt
}

// Decode - decode the array of strings
func Decode(arr []string) string {
	// your code
	fmt.Println("Decoding:", strings.Join(arr, "|"))
	var result string
	shiftVal = -(int(arr[0][1]) - int(arr[0][0]))
	arr[0] = arr[0][2:]
	for _, s := range arr {
		result += strings.Map(mapFn, s)
	}
	return result
}

/////////////////////////////////////////////////////////////////////////////
//package main
//
//import (
//    "strings"
//)
//
//func shifChar(c rune, shift int) string {
//    var shifted rune = 0
//    if c >= 'a' && c <= 'z' {
//        shifted = 'a' + rune((int(c) - 'a' + shift) % 26)
//    } else {
//        if c >= 'A' && c <= 'Z' {
//            shifted = 'A' + rune((int(c) - 'A' + shift) % 26)
//        } else {
//            shifted = c
//        }
//    }
//    return string(shifted)
//}
//func prefix(c rune, shift int) string {
//    var cc = c
//    if c >= 'A' && c <= 'Z' {
//        cc = rune(int(c) + 32)
//    }
//    shifted := 'a' + rune(Modx(int(cc) - 'a' + shift, 26))
//    return string(cc) + string(shifted)
//}
//func Encode(s string, shift int) []string {
//    r := []rune(s)
//    var v = prefix(r[0], shift)
//    for _, c := range r {
//        v += shifChar(c, shift)
//    }
//    var p = len(v) / 5
//    if (len(v) % 5) != 0 {p++}
//    var res = []string{}; var tmp = ""; var cpt = 0
//    for i := 0; i < len(v); i++ {
//        tmp += string(v[i])
//        if ((i + 1) % p == 0) {
//            res = append(res, tmp)
//            tmp = ""
//            cpt++
//        }
//    }
//    if tmp != "" {res = append(res, tmp)}
//    return res
//}
//func Decode(arr []string) string {
//    s := strings.Join(arr, "")
//    r := []rune(s)
//    res := ""
//    shift := int((r[1] - r[0] + 26) % 26)
//    text := r[2:]
//    for _, c := range text {
//        res += shifChar(rune(c), 26 - shift)
//    }
//    return res
//}
