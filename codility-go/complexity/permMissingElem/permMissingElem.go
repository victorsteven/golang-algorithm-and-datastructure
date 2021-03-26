package permMissingElem

import (
	"math"
	"sort"
)

/*
A zero-indexed array A consisting of N different integers is given.
The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

func Solution(A []int) int

that, given a zero-indexed array A, returns the value of the missing element.

For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Assume that:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].

Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).

*/

// Use Sum formula to calculate Sum(len(A)+1) - Sum(A). This will pick up the difference of the 2 sums
func PerMissingElem(A []int) int {
	if len(A) == 0 {
		return 1
	}

	n := len(A) + 1
	arraySum := 0
	totalSum := (n + 1) * n / 2

	for _, value := range A {
		arraySum += value
	}

	return totalSum - arraySum

}

func mine(A []int) int {

	//sort the array:
	sort.Ints(A)

	x := 1

	for i := 0; i < len(A); i++ {
		if x < A[i] {
			return x
		}
		//x = A[i] + 1
		x = int(math.Abs(float64(A[i])) + 1)
	}

	return x

}

//function solution(A) {
//// only positive values, sorted
//A = A.filter(x => x >= 1).sort((a, b) => a - b)
//
//let x = 1
//
//for(let i = 0; i < A.length; i++) {
//// if we find a smaller number no need to continue, cause the array is sorted
//if(x < A[i]) {
//return x
//}
//x = A[i] + 1
//}
//
//return x
//}
