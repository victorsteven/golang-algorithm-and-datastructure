package main

import "fmt"

//Buddy pairs
//You know what divisors of a number are. The divisors of a positive integer n are said to be proper when you consider only the divisors other than n itself. In the following description, divisors will mean proper divisors. For example for 100 they are 1, 2, 4, 5, 10, 20, 25, and 50.
//
//Let s(n) be the sum of these proper divisors of n. Call buddy two positive integers such that the sum of the proper divisors of each number is one more than the other number:
//
//(n, m) are a pair of buddy if s(m) = n + 1 and s(n) = m + 1
//
//For example 48 & 75 is such a pair:
//
//Divisors of 48 are: 1, 2, 3, 4, 6, 8, 12, 16, 24 --> sum: 76 = 75 + 1
//Divisors of 75 are: 1, 3, 5, 15, 25 --> sum: 49 = 48 + 1
//Task
//Given two positive integers start and limit, the function buddy(start, limit) should return the first pair (n m) of buddy pairs such that n (positive integer) is between start (inclusive) and limit (inclusive); m can be greater than limit and has to be greater than n
//
//If there is no buddy pair satisfying the conditions, then return "Nothing" or (for Go lang) nil
//
//Examples
//(depending on the languages)
//
//buddy(10, 50) returns [48, 75]
//buddy(48, 50) returns [48, 75]
//or
//buddy(10, 50) returns "(48 75)"
//buddy(48, 50) returns "(48 75)"
//Note
//for C: The returned string will be free'd.
//See more examples in "Sample Tests:" of your language.

func main() {
	//fmt.Println(getDivisorSum(100))
	//fmt.Println(sumPropDiv(100))

	A := []int{5,3,1,9,2,5,6,3,2}
	fmt.Println(bubble(A))
	fmt.Println(remove_dup(A))
	fmt.Println(largest_num(A))
	fmt.Println(reverse(A)) //this modifies the original array, not the one the duplicate have been removed

	rightAngle(9)

	fmt.Println(fib(9))




	//fmt.Println(array_sort(A))

	//B := []int{5,3,1,9,2,5,6,3,2}
	//fmt.Println("The sort: ", array_sort(B))




}

func buddy_pairs(start, limit int) []int {

	//div

	//var m, n int

	for i := start; i <= limit; i++ {

	}
	return nil
}

func getDivisorSum(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i < n; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	return sum
}


func Buddy(start, limit int) []int {
	for n := start; n <= limit; n++ {
		m := sumPropDiv(n)
		if m > n {
			if sumPropDiv(m) == n {
				return []int{n, m}
			}
		}
	}
	return nil
}

func sumPropDiv(num int) (sum int) {
	for i := 2; i < num/i; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return
}



func bubble(a []int) []int  {
	for i := len(a)-1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if a[j-1] > a[j] {
				temp := a[j-1]
				a[j-1] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func remove_dup(a []int) []int {

	newArray := []int{}
	keys := make(map[int]bool)
	for _, v := range a {
		if value := keys[v]; !value {
			keys[v] = true
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func largest_num(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

//func reverse(a []int) []int  {
//	for i := 0; i < len(a)/2; i++ {
//		j := len(a) - 1 - i
//		a[i], a[j] = a[j], a[i]
//	}
//	return a
//}

func reverse(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func rightAngle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}


func fib(n int) int  {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}