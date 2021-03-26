package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//Given an array of positive or negative integers
//
//I= [i1,..,in]
//
//you have to produce a sorted array P of the form
//
//[ [p, sum of all ij of I for which p is a prime factor (p positive) of ij] ...]
//
//P will be sorted by increasing order of the prime numbers. The final result has to be given as a string in Java, C#, C, C++ and as an array of arrays in other languages.
//
//Example:
//
//I = (/12, 15/); // result = "(2 12)(3 27)(5 15)"
//[2, 3, 5] is the list of all prime factors of the elements of I, hence the result.
//
//Notes:
//
//It can happen that a sum is 0 if some numbers are negative!
//Example: I = [15, 30, -45] 5 divides 15, 30 and (-45) so 5 appears in the result, the sum of the numbers for which 5 is a factor is 0 so we have [5, 0] in the result amongst others.
//
//In Fortran - as in any other language - the returned string is not permitted to contain any redundant trailing whitespace: you can use dynamically allocated character strings.

func main() {

	//A := []int{15,21,24,30,45}

	A := []int{12, 15}
	//lst2 := []int{15,21,24,30,45}

	fmt.Println("The prime: list ", SumOfDivided(A))
	//fmt.Println("The prime: list ", SumOfDivided(lst2))
}

func SumOfDivided(lst []int) string {
	factorMap := make(map[int]int)
	for _, v := range lst {
		for _, w := range unique_primes(v) {
			//provided the key exists, add the "v" values
			factorMap[w] = factorMap[w] + v
		}
	}
	fmt.Println("the factor map: ", factorMap)
	keys := []int{}
	for i := range factorMap {
		keys = append(keys, i)
	}
	sort.Ints(keys) //sort the array
	result := ""
	for _, v := range keys {
		result += "(" + strconv.Itoa(v) + " " + strconv.Itoa(factorMap[v]) + ")"
	}
	return result
}

func unique_primes(n int) []int {
	absNum := math.Abs(float64(n)) //change the negatives to positive
	numInt := int(absNum)
	divisor := 2
	primes := []int{}
	for numInt > 1 {
		if numInt%divisor == 0 {
			primes = append(primes, divisor)
			numInt /= divisor
		} else {
			divisor++
		}
	}
	//get only the unique value from the prime array:
	keys := map[int]bool{}
	unique_prime := []int{}
	for _, v := range primes {
		if value := keys[v]; !value {
			keys[v] = true
			unique_prime = append(unique_prime, v)
		}
	}
	return unique_prime
}

//func sumElements(n []int) int {
//	ans := 0
//	for _, v := range n {
//		ans += v
//	}
//	return ans
//}
