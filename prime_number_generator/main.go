package main

import "fmt"

//Your task here is to generate a list of prime numbers, starting from 2.
//
//One way of doing this is using python's generators.
//
//In case you choose to use generators, notice that the generator object should contain all the generated prime numbers, from 2 to an upper limit (included) provided as the argument to the function. If the input limit is less than 2 (including negatives), it should return an empty list.
//
//Each iteration of the generator will yield one prime number. See this link for reference.
//
//The generator object will be converted to a list outside the code, within the test cases.
//
//There will be no check if you are using generators or lists, so use the one you feel more confortable with.

func main() {
	fmt.Println("this is the one: ", list(10))
}

func list(n int) []int {

	arr := []int{}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
