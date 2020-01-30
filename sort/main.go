package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,8,3,9,2}
	fmt.Println(sortInt(a))

	str := []string{"hello man", "steven", "this", "is a good piece"}
	fmt.Println(sortString(str))
}


func sortInt(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

//sort the sentence according to the length of the string
func sortString(str []string) []string {
	sort.Slice(str , func(i, j int) bool {
		return len(str[i]) < len(str[j])
	})
	return str
}

func sortString2(str []string) []string {
	sort.Slice(str, func(i, j int) bool {
		return len(str[i]) < len(str[j])
	})
}




