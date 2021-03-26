package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(10))

	str := "This is the awesome birthday here. Okay"
	fmt.Println(vowel(str))

}

func sum(n int) float64 {
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * i
	}
	//the mean:
	mean := math.Sqrt(float64(ans / n))

	return mean
}

//count the vowels in a word

func vowel(str string) int {
	count := 0
	for _, v := range str {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
