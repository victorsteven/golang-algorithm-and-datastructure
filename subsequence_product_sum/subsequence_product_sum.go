package main

//package subsequence_product_sum

import "fmt"

//Given an array of integers and an integer m, return the sum of the product of its subsequences of length m.
//
//Ex1:
//
//a := []int{1, 2, 3}
//m := 2
//the subsequences(of length 2) are (1,2) (1,3) (2,3), you should multiply the numbers of each subsequence and take their sum
//
//ProductSum = (1*2) + (1*3) + (2*3) //=> 11
//Ex2:
//
//a := []int{2,3,4,5}
//m := 3
//the subsequences(of length 3) are (2,3,4) (2,4,5) (2,3,5) (3,4,5)
//
//ProductSum = (2*3*4) + (2*3*5) + (2*4*5) + (3*4*5) //=> 154
//Task:
//Write a function productSum(a, m) that does as described above
//
//The sum can be really large, so return the result in modulo 109+7
//
//Constraints
//0 <= A[i] < 1000000
//
//1 < k <= 50
//99 random tests |A| = 10^5
//1 big test |A| = 10^6
//k < |A|

func main() {
	a := ProductSum([]int{1, 2, 3}, 2)

	b := ProductSum([]int{2, 3, 4, 5}, 3)

	fmt.Println(a)
	fmt.Println(b)
}

//func ProductSum(a []int, m int) int {
//
//	//result := 0
//
//	for i := 0; i < len(a); i++ {
//
//		rel := a[i:m]
//
//		fmt.Println(rel)
//
//		//result =
//
//	}
//
//	return -1
//}

func ProductSum(a []int, m int) int {
	sm, n, M := make([]int, len(a)), len(a), 1000000007

	p := make([][]int, n+1)

	copy(sm, a)

	for i := 0; i <= n; i++ {
		p[i] = make([]int, m+1)
	}
	for i := n - 2; i >= 0; i-- {
		sm[i] = (sm[i] + sm[i+1]) % M
	}
	for i := n - 1; i >= 0; i-- {
		for k := 1; k < m+1; k++ {
			if k == 1 {
				p[i][k] = sm[i]
				continue
			}
			p[i][k] = ((a[i] * p[i+1][k-1] % M) + p[i+1][k]) % M
		}
	}
	return p[0][m]
}
