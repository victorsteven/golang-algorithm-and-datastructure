package main

import "fmt"

func main() {

	A := []int{2, 3, 4, 6, 8, 10}

	fmt.Println("this is the code: ", FindOutlier(A))
}

func FindOutlier(integers []int) int {
	outlier := 0
	odd_array := []int{}
	even_array := []int{}
	for _, v := range integers {
		if v%2 == 0 {
			even_array = append(even_array, v)
		} else {
			odd_array = append(odd_array, v)
		}
	}
	if len(odd_array) > len(even_array) {
		outlier = even_array[0]
	} else {
		outlier = odd_array[0]
	}
	return outlier
}
