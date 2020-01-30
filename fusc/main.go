package main

import "fmt"

//The fusc function is defined recursively as follows:
//
//1. fusc(0) = 0
//2. fusc(1) = 1
//3. fusc(2 * n) = fusc(n)
//4. fusc(2 * n + 1) = fusc(n) + fusc(n + 1)
//The 4 rules above are sufficient to determine the value of fusc for any non-negative input n. For example, let's say you want to compute fusc(10).
//
//fusc(10) = fusc(5), by rule 3.
//fusc(5) = fusc(2) + fusc(3), by rule 4.
//fusc(2) = fusc(1), by rule 3.
//fusc(1) = 1, by rule 2.
//fusc(3) = fusc(1) + fusc(2) by rule 4.
//fusc(1) and fusc(2) have already been computed are both equal to 1.
//Putting these results together fusc(10) = fusc(5) = fusc(2) + fusc(3) = 1 + 2 = 3
//
//Your job is to produce the code for the fusc function. In this kata, your function will be tested with small values of n, so you should not need to be concerned about stack overflow or timeouts.
//
//Hint: Use recursion.
//
//When done, move on to Part 2.

func main() {
	fmt.Println(fusc(9))
	fmt.Println(fusc2(2))
	fmt.Println(fusc3(4))

	//fmt.Println(genFib(9))
	//fmt.Println(Fibbonacci([2]float64{1,9}, 2))
	//fmt.Println(Fibbonacci([2]float64{1,1}, 10))
}

func fusc(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n % 2 == 0 {
		return fusc(n/2)
	}
	n = (n-1)/2
	return fusc(n) + fusc(n+1)
}

func fusc2(n int) int {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	case n%2 == 0:
		return fusc2(n / 2)
	default:
		n = (n-1) / 2
		return fusc2(n) + fusc2(n+1)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}


func Fibbonacci(signature [2]float64, n int) []float64 {
	result := signature[:]
	if n <= 2 {
		return result[:n]
	}
	for i := 2; i < n; i++ {
		sum := result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

//generating fib:
func Tribonacci3(signature [3]float64, n int) []float64 {
	result := signature[:]
	if n <= 3 {
		return result[:n]
	}
	for i := 3; i < n; i++ {
		sum := result[i-3] + result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

func genFib(n int) []float64 {
	result := []float64{1,1}
	if n < 2 {
		return result[:n]
	}
	for i := 2; i < n; i++ {
		//fmt.Println("this is i: ", i)
		//fmt.Println(result[i-1])
		sum := result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

func fusc3(n int) int {
	if n < 2 {
		return n
	}
	if n % 2 == 0 {
		return fusc3(n/2)
	}
	n = (n-1)/2
	return fusc3(n) + fusc3(n+1)
}
