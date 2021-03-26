package main

import "fmt"

func Snail(snailMap [][]int) []int {
	x := len(snailMap[0])
	myLen := x * x
	res := []int{}

	round := 0

	for myLen > 0 {

		for i := round; i < x-round; i++ {
			res = append(res, snailMap[round][i])
			myLen -= 1
		}

		for i := round + 1; i < x-round; i++ {
			res = append(res, snailMap[i][x-1-round])
			myLen -= 1
		}

		for i := x - 2 - round; i >= round; i-- {
			res = append(res, snailMap[x-1-round][i])
			myLen -= 1
		}

		for i := x - 2 - round; i > round; i-- {
			res = append(res, snailMap[i][round])
			myLen -= 1
		}
		round += 1
	}
	return res
}

func main() {

	array1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans1 := Snail(array1)
	fmt.Println(ans1)
}
