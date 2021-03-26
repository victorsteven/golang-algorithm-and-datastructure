package main

import (
	"fmt"
	"sort"
)

func main() {
	var ls = []int{50, 55, 57, 58, 60}
	var k = 3
	var t = 174

	fmt.Println(ChooseBestSum(t, k, ls))
}

//func ChooseBestSum(t, k int, ls []int) int {
//	if len(ls) < k {
//		return -1
//	}
//	tripCounter := 0
////	sorting the array:
//	totalDistance := 0
//	sort.Ints(ls)
//
//	sort.Slice(ls, func(i, j int) bool {
//		return ls[i] > ls[j]
//	})
//	for _, v := range ls {
//		if v < t && (v + totalDistance) < t && tripCounter <= k {
//			tripCounter++
//			totalDistance += v
//		}
//	}
//	return totalDistance
//}

func ChooseBestSum(t, k int, ls []int) int {
	outerbest := -1
	for i, d := range ls {
		//not enough remaining values for this d to work
		if len(ls) < k {
			continue
		}
		//recursively choose best from t-d, until final level k = 1
		if k > 1 {
			innerbest := ChooseBestSum(t-d, k-1, ls[i+1:])
			//if no best available at lower level, this d cant work
			if innerbest < 0 {
				continue
			}
			d += innerbest
		}
		if d <= t && d > outerbest {
			outerbest = d
		}
	}
	return outerbest
}

var buffer [][]int

func ChosseBestSum2(t, k int, ls []int) int {
	buffer = buffer[0:0]
	data := make([]int, k)
	n := len(ls)
	sort.Ints(ls)
	combinations(ls, data, 0, n-1, 0, k)
	max := -1
	sum := 0
	for _, value := range buffer {
		for _, v := range value {
			sum += v
		}
		if sum <= t && max < sum {
			max = sum
		}
		sum = 0
	}
	return max
}

func combinations(arr, data []int, start, end, index, r int) {
	if index == r {
		var i = make([]int, 0)
		i = append(i, data...)
		buffer = append(buffer, i)
		return
	}
	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = arr[i]
		combinations(arr, data, i+1, end, index+1, r)
	}
}
