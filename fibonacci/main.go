package main

import "fmt"

func main() {
	//fmt.Println(fib(6))

	//using channel:
	//c := make(chan int)
	//go fibGen(c)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c)
	//}

	//Generating fib without channels:
	//for i := 0; i < 10; i++ {
	//	//	fmt.Print(fibGen2(i), " ")
	//	//}

	//fmt.Println(FibGen3([2]float64{1,9}, 10))
	fmt.Println(FibGenP([2]float64{1, 9}, 10))
	fmt.Println(FibGenRange(10))

	//arr := []int{1,2,3,4,5,6,7}
	//arr = arr[1:4]
	//fmt.Println("this is a training: ", arr)
}

//Write a function that returns a particular fib number
//Using Recursion
//func fib(a int) int {
//	if a < 2 {
//		return a
//	}
//	return fib(a - 1) + fib(a - 2)
//}

//Write a function that generates a fib sequence
//using channels and goroutines
func fibGen(c chan int) {
	a, b := 0, 1
	for {
		c <- a
		a, b = b, a+b
	}
}

//Generating Fibonacci without channels:
func fibGen2(a int) int {
	//if a < 2 {
	//	return a
	//}
	//for the sake of negative fib:
	if a == 0 || a == 1 {
		return a
	}
	return fibGen2(a-1) + fibGen2(a-2)
}

func fibo(a int) int {
	if a == 0 || a == 1 {
		return a
	}
	return fibo(a-1) + fibo(a-2)
}

func Fibonacci(signature [2]float64, n int) (r []float64) {
	r = signature[:] //since we told the array the number of elements to expect, we use [:]
	for i := 0; i < n; i++ {
		l := len(r)
		r = append(r, r[l-1]+r[l-2])
	}
	return r[:n]
}

//We didnt tell the array the number of elements to expect, but from the implementation the array should have just two elements
//func Fibonacci(signature []float64, n int) (r []float64) {
//	r = signature
//	for i := 0; i < n; i++ {
//		l := len(r)
//		r = append(r, r[l-1] + r[l-2])
//	}
//	return r[:n]
//}

//more readable
func FibGen3(signature [2]float64, n int) []float64 {
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

func FibGenP(sig [2]float64, n int) []float64 {
	result := sig[:]
	if n == 0 || n == 1 {
		return result[:n]
	}
	for i := 2; i < n; i++ {
		sum := result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

func FibGenRange(n int) []float64 {
	result := []float64{1: 1}
	if n == 0 || n == 1 {
		return result[:n]
	}
	for i := 2; i < n; i++ {
		sum := result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}
