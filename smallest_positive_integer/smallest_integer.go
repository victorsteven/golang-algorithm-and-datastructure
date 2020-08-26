package smallest_positive_integer

import "sort"

//This is a demo task.
//
//Write a function:
//
//func Solution(A []int) int
//
//that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
//
//For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.
//
//Given A = [1, 2, 3], the function should return 4.
//
//Given A = [−1, −3], the function should return 1.
//
//Write an efficient algorithm for the following assumptions:
//
//N is an integer within the range [1..100,000];
//each element of array A is an integer within the range [−1,000,000..1,000,000].
//Copyright 2009–2020 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


func Solution(A []int) int {

	//sort the array:
	sort.Ints(A)

	//let the smallest positive integer be 1
	x := 1

	for i := 0; i < len(A); i++ {
		if A[i] == x {
			x++
		}
	}

	return x

}
