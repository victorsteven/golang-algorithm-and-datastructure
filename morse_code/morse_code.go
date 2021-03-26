package morse_code

import (
	. "math"
	"regexp"
	"strings"
)

func getUnitLen(t string) int {
	r := regexp.MustCompile(`1+|0+`)
	allSubseq := r.FindAllString(t, -1)
	var minLen int
	for i, v := range allSubseq {
		if i == 0 || minLen > len(v) {
			minLen = len(v)
		}
	}
	return minLen
}

func DecodeBits(bits string) (res string) {
	b := strings.Trim(bits, "0")
	if len(b) == 0 {
		return " "
	}
	unit := getUnitLen(b)
	pause := strings.Repeat("0", unit)
	c := strings.Split(b, pause)
	dot := strings.Repeat("1", unit)
	dash := strings.Repeat(dot, 3)
	var space int
	for _, v := range c {
		if len(v) == 0 {
			space++
		}
		if space == 2 {
			res += " "
			space = 0
		}
		if v == dot {
			res += "."
		}
		if v == dash {
			res += "-"
		}
	}
	return res
}

func DecodeMorse(morseCode string) (res string) {
	m := strings.TrimSpace(morseCode)
	if len(m) == 0 {
		return res
	}
	c := strings.Split(m, " ")
	var space int
	for _, v := range c {
		if len(v) == 0 {
			space++
		}
		if space == 2 {
			res += " "
			space = 0
		}
		res += MORSE_CODE[v]
	}
	return res
}

func DecodeBits2(bits string) string {
	bits = compact(strings.Trim(bits, "0"))

	bits = replace(bits, `111(0|$)`, "-")
	bits = replace(bits, `1(0|$)`, ".")
	bits = replace(bits, `0`, " ")

	return bits
}

func DecodeMorse2(morseCode string) string {
	decoded := ""

	for _, word := range strings.Split(morseCode, "   ") {
		decodedWord := ""

		for _, char := range strings.Split(word, " ") {
			decodedWord += MORSE_CODE[char]
		}

		if decodedWord != "" {
			decoded += " " + decodedWord
		}
	}

	return strings.Trim(decoded, " ")
}

func replace(src, re, newSubStr string) string {
	return regexp.MustCompile(re).ReplaceAllString(src, newSubStr)
}

func compact(bits string) string {
	chunks := regexp.MustCompile(`(0+|1+)`).FindAllString(bits, -1)

	// Find minimum chunk length
	min := MaxInt16
	for _, str := range chunks {
		if len(str) < min {
			min = len(str)
		}
	}

	// Compact string
	compacted := ""
	for _, str := range chunks {
		compacted += str[0:(len(str) / min)]
	}

	return compacted
}
