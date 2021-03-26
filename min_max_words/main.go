package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("this is the ans low: ", FindLow("normal_http_call my name is a"))
	fmt.Println("this is the ans high: ", FindHigh("normal_http_call my name is a"))

	low, high := FindLowHigh("normal_http_call my name is a")
	fmt.Println("Lowest: ", low)
	fmt.Println("Highest: ", high)
}

func FindLow(s string) int {
	if len(s) <= 0 {
		return 0
	}
	words := strings.Split(s, " ")

	lowestWord := words[0]
	for _, v := range words {
		if len(v) < len(lowestWord) {
			lowestWord = v
		}
	}
	return len(lowestWord)
}

func FindHigh(s string) int {
	if len(s) <= 0 {
		return 0
	}
	words := strings.Split(s, " ")

	lowestWord := words[0]
	for _, v := range words {
		if len(v) > len(lowestWord) {
			lowestWord = v
		}
	}
	return len(lowestWord)
}

func FindLowHigh(s string) (int, int) {
	if len(s) <= 0 {
		return 0, 0
	}
	words := strings.Split(s, " ")

	lowestWord := words[0]
	highestWord := words[0]

	for _, v := range words {
		if len(v) < len(lowestWord) {
			lowestWord = v
		}
	}

	for _, v := range words {
		if len(v) > len(highestWord) {
			highestWord = v
		}
	}
	return len(lowestWord), len(highestWord)
}
