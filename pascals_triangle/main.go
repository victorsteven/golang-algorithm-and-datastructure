package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println()
	pascal(10)
}

func binomialCoeff(n, k int) int {
	res := 1
	if k > n-k {
		k = n - k
	}
	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}

func pascal(n int) {
	for line := 0; line < n; line++ {
		for i := 0; i <= line; i++ {
			fmt.Print("" + strconv.Itoa(binomialCoeff(line, i)) + " ")
		}
		fmt.Println()
	}
}
