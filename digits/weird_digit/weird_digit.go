package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Some numbers have funny properties. For example:
//
//89 --> 8¹ + 9² = 89 * 1
//
//695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2
//
//46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
//
//Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p
//
//we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.
//In other words:
//
//Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k
//
//If it is the case we will return k, if not return -1.
//
//Note: n and p will always be given as strictly positive integers.

func DigPow(n, p int) int {

	str := strconv.Itoa(n)
	strResult := strings.Split(str, "")

	numArr := make([]int, 0)

	for _, v := range strResult {
		num, _ := strconv.Atoi(v)
		numArr = append(numArr, num)
	}

	arrSum := 0
	length := len(str)
	for i := 0; i < length; i++ {
		arrSum += int(math.Pow(float64(numArr[i]), float64(p)))
		p++
	}

	if arrSum%n == 0 {
		return arrSum / n
	} else {
		return -1
	}
}

func main() {
	fmt.Println(DigPow(46288, 3))
	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(92, 1))
	fmt.Println(DigPow(695, 2))
}

//func DigPowa(n, p int) int {
//	var sum int = 0
//
//	nstring := strconv.Itoa(n)
//
//	for i := 0; i < len(nstring); i, p = i+1, p+1 {
//		digit, _ := strconv.Atoi(string(nstring[i]))
//		sum += int(math.Pow(float64(digit), float64(p)))
//	}
//
//	rest := sum % n
//	if rest == 0 {
//		return sum / n
//	} else {
//		return -1
//	}
//}

//func DigPowb(n, p int) int {
//	N := n
//	var digits []int;
//
//	for n > 0 {
//		digits = append(digits, n % 10)
//		n = n / 10
//	}
//
//	sum := 0
//	for i := len(digits) - 1; i >= 0; i-- {
//		sum = sum + int(math.Pow(float64(digits[i]), float64(p)))
//		p = p + 1
//	}
//
//	if sum % N == 0 {
//		return sum / N
//	} else {
//		return -1
//	}
//}
