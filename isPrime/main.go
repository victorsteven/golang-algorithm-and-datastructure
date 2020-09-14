package main

import "fmt"

func main() {
	fmt.Println(isPrime(6))
	fmt.Println(isPrime(7))
	fmt.Println(isPrime(9))
	fmt.Println(isPrime(2))
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}