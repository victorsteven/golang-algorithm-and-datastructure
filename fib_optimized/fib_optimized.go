package main

import (
	"fmt"
	"log"
	"time"
)

//This is from the work of @adeshinaHH

func main() {

	fmt.Println("the main started")

	input := int64(10)

	//recursion
	start1 := time.Now()
	result1 := expensiveFib(input)
	defer trackTime(start1, result1, "Recursive for first function")
	fmt.Println()

	//forloop
	start2 := time.Now()
	result2 := FibByLoop(input)
	defer trackTime(start2, result2, "Forloop Implementation")
	fmt.Println()

	//memoization
	start3 := time.Now()
	val := memoize(expensiveFib)
	result3 := val(input)
	defer trackTime(start3, result3, "Memoization")

}

func expensiveFib(n int64) int64 {

	if n < 2 {
		return n
	}

	result := expensiveFib(n-1) + expensiveFib(n-2)

	return result
}

func FibByLoop(n int64) int64 {

	fibBox := []int64{0, 1}

	for i := int64(0); i < n; i++ {

		value := fibBox[i] + fibBox[i + 1]

		fibBox = append(fibBox, value)
	}

	result := fibBox[n]

	return result
}

func trackTime(start time.Time, result int64, name string) {

	elapsed := time.Since(start)

	log.Printf("----> %s solution | result: %v | took: %s", name, result, elapsed)
}

type funcFib func(int64)int64

func memoize(fn funcFib) func(int64)int64 {
	cache := make(map[int64]int64, 0)

	return func(n int64) int64 {
		if _, ok := cache[n]; ok {
			return cache[n]
		}
		result := fn(n)
		cache[n] = result

		return result
	}
}