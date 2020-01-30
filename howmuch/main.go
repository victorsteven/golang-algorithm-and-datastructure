package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("this is the result: ", HowMuch(1, 100))
}

func HowMuch(m int, n int) [][3]string {
	var results [][3]string

	if m > n {
		m, n = n, m
	}

	for i := m; i <= n; i++ {
		if (i - 1) % 9 == 0 && (i - 2) % 7 == 0 {
			results = append(results, [3]string{
				"M: " + strconv.Itoa(i),
				"B: " + strconv.Itoa((i - 2) / 7),
				"C: " + strconv.Itoa((i - 1) / 9),
			})
		}
	}
	return results
}























