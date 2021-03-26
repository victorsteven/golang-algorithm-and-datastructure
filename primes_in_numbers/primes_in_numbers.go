package primes_in_numbers

import "fmt"

//Given a positive number n > 1 find the prime factor decomposition of n. The result will be a string with the following form :
//
// "(p1**n1)(p2**n2)...(pk**nk)"

//where a ** b means a to the power of b
//with the p(i) in increasing order and n(i) empty if n(i) is 1.
//Example: n = 86240 should return "(2**5)(5)(7**2)(11)"

//func PrimeFactors(n int) string {
//
//	primes := getPrimes(n)
//
//	hashMap := make(map[int]int)
//
//	result := ""
//	for _, v := range primes {
//		hashMap[v]++
//	}
//
//	fmt.Println("the hash: ", hashMap)
//
//	for i, v := range hashMap {
//
//		fmt.Println("each: ", i)
//
//		if v  == 1 {
//			result += fmt.Sprintf("(%d)", i)
//		} else if v > 1 {
//			result += fmt.Sprintf("(%d**%d)", i, v)
//		}
//	}
//
//	fmt.Println("the result now: ", result)
//
//	return result
//
//}
//
//func getPrimes(n int) (primes []int) {
//
//	div := 2
//
//	for n > 1 {
//		if n % div == 0 {
//			primes = append(primes, div)
//			n = n / div
//		} else {
//			div++
//		}
//	}
//	return primes
//}

func PrimeFactors(n int) (ret string) {
	for i, m := 2, 0; n > 0; {
		if n%i != 0 {
			if m == 1 {
				ret += fmt.Sprintf("(%d)", i)
			} else if m > 1 {
				ret += fmt.Sprintf("(%d**%d)", i, m)
			}
			if n == 1 {
				break
			}
			i, m = i+1, 0
			continue
		}
		n, m = n/i, m+1
	}

	fmt.Println("ans: ", ret)

	return ret
}
