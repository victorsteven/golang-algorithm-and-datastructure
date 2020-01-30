package main

import "fmt"

func main() {
	A := []int{20,2,16,4,6,7,34,2,9,15,5}
	fmt.Println(max_num(A))
}


func max_num(arr []int) int {

	maxnumber := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxnumber {
			maxnumber = arr[i]
		}
	}
	return maxnumber
}