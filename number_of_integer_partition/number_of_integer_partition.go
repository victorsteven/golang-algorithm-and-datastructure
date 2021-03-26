package number_of_integer_partition

import "fmt"

//An integer partition of n is a weakly decreasing list of positive integers which sum to n.
//
//For example, there are 7 integer partitions of 5:
//
//[5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1].
//
//Write a function named partitions which returns the number of integer partitions of n. The function should be able to find the number of integer partitions of n for n as least as large as 100.

func p(n, m int) int {

	if n <= 1 {
		return 1
	}
	if m > n {
		return p(n, n)
	}

	sum := 0
	for k := 1; k <= m; k++ {
		fmt.Println("ans passed: ", n-k, k)
		sum += p(n-k, k)
		fmt.Println("sum: ", sum)
	}
	return sum
}

func Partitions(n int) int {
	return p(n, n)
}
