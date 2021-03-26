package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{2, -7, -2, -2, 0}
	fmt.Println(absSort(A))
}

func absSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if math.Abs(float64(arr[j-1])) > math.Abs(float64(arr[j])) {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
			if math.Abs(float64(arr[j-1])) == math.Abs(float64(arr[j])) && arr[j-1] > 0 {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
