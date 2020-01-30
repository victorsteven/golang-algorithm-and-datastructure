package main

import "fmt"

func main() {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("max sum: ", max_subarray(A))
	fmt.Println("max sum: ", max_subarray_2(A))

	//fmt.Println(max_subarray_plus(A))
	//fmt.Println()
}

//The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers:
//maxSequence [-2, 1, -3, 4, -1, 2, 1, -5, 4]
//-- should be 6: [4, -1, 2, 1]

//Easy case is when the list is made up of only positive numbers and the maximum sum is the sum of the whole array. If the list is made up of only negative numbers, return 0 instead.

//Empty list is considered to have zero greatest sum. Note that the empty list or array is also a valid sublist/subarray.

func max_subarray(numbers []int) int {
	var best, sum int
	for _, x := range numbers {
		sum += x
		//fmt.Printf("The current %d and the sum %d \n", x, sum)
		//fmt.Println(sum)
		switch {
		case sum > best:
			best = sum
		case sum < 0:
			sum = 0
		}
	}
	return best
}

//Another approach
func max_subarray_2(numbers []int) int {
	max, curr := 0, 0
	for _, v := range numbers {
		if curr += v; curr < 0 {
			curr = 0
		} else if curr > max {
			max = curr
		}
	}
	return max
}




//Return both the filtered array
func max_subarray_plus(numbers []int) ([]int, int) {

	var best, start, end, sum, sumStart int
	for i, x := range numbers {
		sum += x
		switch {
		case sum > best:
			best = sum
			start = sumStart
			end = i + 1
		case sum < 0:
			sum = 0
			sumStart = i + 1
		}
	}
	return numbers[start:end], best
}