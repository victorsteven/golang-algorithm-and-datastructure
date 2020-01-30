package main

import (
	"fmt"
	"sort"
)

func main() {
	ls := []int{12, 34, 79}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] > ls[j]
	})
	fmt.Println(ls)
}