package main

import "fmt"

func main() {
	fmt.Println("the fac: ", fac(5))
	fmt.Println("the fib number: ", fib(7))

	c := make(chan int)
	go fibSeries(c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

//Write a function to return the factorial of a number.
//Solutions can be iterative or recursive.
//Recursion solution
func fac(n int) int {
	if n < 2 {
		return 1
	}
	return n * fac(n-1)
}

func fib(a int) int {
	if a < 2 {
		return a
	}
	return fib(a-1) + fib(a-2)
}

func fibSeries(c chan int) {
	a, b := 0, 1
	for {
		c <- a
		a, b = b, a+b
	}
}
