package main

import (
	"fmt"
	"strconv"
	"strings"
)

//In this kata, your task is to implement what I call Iterative Rotation Cipher (IRC). To complete the task, you will create an object with two methods, encode and decode. (For non-JavaScript versions, you only need to write the two functions without the enclosing dict)
//
//Input
//The encode method will receive two arguments — a positive integer n and a string value.
//
//The decode method will receive one argument — a string value.
//
//Output
//Each method will return a string value.
//
//How It Works
//Encoding and decoding are done by performing a series of character and substring rotations on a string input.
//
//Encoding: The number of rotations is determined by the value of n. The sequence of rotations is applied in the following order:
// Step 1: remove all spaces in the string (but remember their positions)
// Step 2: shift the order of characters in the new string to the right by n characters
// Step 3: put the spaces back in their original positions
// Step 4: shift the characters of each substring (separated by one or more consecutive spaces) to the right by n
//Repeat this process until it has been completed n times in total.
//The value n is then prepended to the resulting string with a space.
//
//Decoding: Decoding simply reverses the encoding process.

func main() {
	quote := `If you wish to make an apple pie from scratch, you must first invent the universe.`;
    encoded := Encode(10,quote)
	fmt.Println(encoded)

    decoded := Decode(encoded)
	fmt.Println(decoded)

}

func Encode(n int, s string) string {
	i := 0
	spaces := []int{}
	for i < n {
		s, spaces = removeSpaces(s)
		//fmt.Println("nospaces: " + s + "\n")
		//fmt.Println(spaces)
		s = stringShift(s, n)
		//fmt.Println("stringShift: " + s + "\n")
		s = addSpaces(s, spaces)
		//fmt.Println("addSpaces: " + s + "\n")
		s = wordShift(s, n, spaces)
		//fmt.Println("words shifted: " + s + "\n")
		i++
	}
	strN := strconv.Itoa(n)
	return strN + " " + s
}

func Decode(s string) string {
	i := 0
	space := strings.Index(s, " ")
	n, _ := strconv.Atoi(s[:space])
	//fmt.Print("n: ")
	//fmt.Println(n)
	words := s[(space + 1):]
	//fmt.Println("encoded string: " + words)
	for i < n {
		_, spaces := removeSpaces(words)
		words = wordShift(words, -n, spaces)
		//fmt.Println("wordShift: " + words)
		words, _ = removeSpaces(words)
		//fmt.Println("spaces removed: " + words)
		words = stringShift(words, -n)
		//fmt.Println("string shifted: " + words)
		words = addSpaces(words, spaces)
		//fmt.Println("spaces added again" + words)
		i++
	}
	return words
}

func removeSpaces(s string) (string, []int) {
	spaces := []int{}
	for i, c := range s {
		if c == ' ' {
			spaces = append(spaces, i)
		}
	}
	noSpaces := strings.Replace(s, " ", "", -1)
	return noSpaces, spaces
}

func addSpaces(s string, spaces []int) string {
	firstN := ""
	ending := ""
	sOut := s
	for _, space := range spaces {
		firstN = string(sOut[:space])
		ending = string(sOut[space:])
		sOut = firstN + " " + ending
	}
	return sOut
}

func stringShift(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	if n > 0 {
		n = n % len(s)
		beginning := string(s[:len(s)-n])
		lastN := string(s[len(s)-n:])
		return lastN + beginning
	} else if n < 0 {
		n = -n
		n = n % len(s)
		firstN := string(s[:n])
		ending := string(s[n:])
		return ending + firstN
	}
	return s
}

func wordShift(s string, n int, spaces []int) string {
	output := ""
	for _, word := range strings.Split(s, " ") {
		output = output + stringShift(word, n)
	}
	return addSpaces(output, spaces)
}