package main

import "fmt"

func main() {
	fmt.Println(unique([]int{1,2,3,5,5,3,2,5}))
	fmt.Println(removeDup([]int{1,2,3,5,5,3,2,5}))

}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeDup(intSlice []int) []int  {
	var list []int
	keys := make(map[int]bool)
	for _, v := range intSlice {
		if _, value := keys[v]; !value { //it is not a duplicate
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}


















