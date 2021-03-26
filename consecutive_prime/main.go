package main

import (
	"fmt"
)

//A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity. A natural number is thus prime if and only if it is 1-prime.
//Examples:
//k = 2 -> 4, 6, 9, 10, 14, 15, 21, 22, …
//k = 3 -> 8, 12, 18, 20, 27, 28, 30, …
//k = 5 -> 32, 48, 72, 80, 108, 112, …

//Task:
//Given an integer k and a list arr of positive integers the function consec_kprimes (or its variants in other languages) returns how many times in the sequence arr numbers come up twice in a row with exactly k prime factors?

//arr = [10005, 10030, 10026, 10008, 10016, 10028, 10004]
//consec_kprimes(4, arr) => 3 because 10005 and 10030 are consecutive 4-primes, 10030 and 10026 too as well as 10028 and 10004 but 10008 and 10016 are 6-primes.
//
//consec_kprimes(4, [10175, 10185, 10180, 10197]) => 3 because 10175-10185 and 10185- 10180 and 10180-10197 are all consecutive 4-primes.

func main() {
	//fmt.Println("this is the prime factors: ", prime_factors(2501))
	//
	//A := []int{10005, 10030, 10026, 10008, 10016, 10028, 10004}
	//fmt.Println("Number of consecutives: ", consecPrime(4, A))

	//ans: 2
	fmt.Println(consecPrime(2, []int{10081, 10071, 10077, 10065, 10060, 10070, 10086, 10083, 10078, 10076, 10089, 10085, 10063, 10074, 10068, 10073, 10072, 10075}))

	//ans: 0
	fmt.Println(consecPrime(6, []int{10064}))
	//
	////ans 0
	fmt.Println(consecPrime(1, []int{10054, 10039, 10053, 10051, 10047, 10043, 10037, 10034}))

	//ans 5
	fmt.Println(consecPrime(3, []int{10158, 10182, 10184, 10172, 10179, 10168, 10156, 10165, 10155, 10161, 10178, 10170}))

	//ans: 7
	fmt.Println(consecPrime(2, []int{10110, 10102, 10097, 10113, 10123, 10109, 10118, 10119, 10117, 10115, 10101, 10121, 10122}))

	//ans 7
	fmt.Println(consecPrime(2, []int{4, 6, 9, 10, 14, 15, 21, 22}))

	//ans 6
	fmt.Println(consecPrime(3, []int{8, 12, 18, 20, 27, 28, 30}))

	//ans 5
	fmt.Println(consecPrime(5, []int{32, 48, 72, 80, 108, 112}))

	//ans 3
	fmt.Println(consecPrime(4, []int{10005, 10030, 10026, 10008, 10016, 10028, 10004}))

	//ans 3
	fmt.Println(consecPrime(4, []int{10175, 10185, 10180, 10197}))
}

func prime_factors(a int) []int {
	divisor := 2
	factors := []int{}
	for a > 1 {
		if a%divisor == 0 {
			factors = append(factors, divisor)
			a = a / divisor
		} else {
			divisor++
		}
	}
	return factors
}

//Another approach
func consecPrime(k int, arr []int) int {
	prime_count := []int{}
	for i := 0; i < len(arr); i++ {
		count := 0
		divisor := 2
		for arr[i] > 1 {
			if arr[i]%divisor == 0 {
				count++
				arr[i] = arr[i] / divisor
			} else {
				divisor++
			}
		}
		prime_count = append(prime_count, count)
	}
	result := 0
	for i := 0; i < len(prime_count)-1; i++ {
		if prime_count[i] == prime_count[i+1] && prime_count[i] == k {
			result++
		}
	}
	return result
}

//func consecPrime2(k int, arr []int) int {
//	//loop through array and find the prime numbers
//	prime_2D := [][]int{}
//	for i := 0; i < len(arr); i++ {
//		each_number_prime := []int{}
//		divisor := 2
//		for arr[i] > 1 {
//			if  arr[i] % divisor == 0 {
//				each_number_prime = append(each_number_prime, divisor)
//				arr[i] = arr[i] / divisor
//			} else {
//				divisor++
//			}
//		}
//		prime_2D = append(prime_2D, each_number_prime)
//	}
//	result := 0
//
//	for i := 0; i < len(prime_2D)-1; i++ {
//		if len(prime_2D[i]) == len(prime_2D[i+1]) &&  len(prime_2D[i]) == k {
//				result++
//		}
//	}
//	return result
//}

//checking for prime number
//func isPrime(a int64) bool {
//	if big.NewInt(a).ProbablyPrime(0) {
//		return true
//	}
//	return false
//}
