package main

import (
	"fmt"
	"strconv"
)

var lst = [][]int{{1, 2}, {1, 3}, {1, 4}}



func main() {
	fmt.Println("this is the statement")
	fmt.Println(ConvertFracts(lst))
}

func ConvertFracts(a [][]int) string  {
	var lcm = 1
	for _, v := range a {
		var denom = v[1] / GCF(v[0], v[1])
		var gcf = GCF(lcm, denom)
		lcm = lcm * (denom / gcf)
	}
	var answer string
	for _, item := range a {
		answer = answer + "(" + strconv.Itoa(item[0] * lcm / item[1]) + "," + strconv.Itoa(lcm) + ")"
	}
	return answer
}

func GCF(a int, b int) int {
	var x, y int
	if a > b {
		x, y = a, b
	} else {
		x, y = b, a
	}
	for y != 0 {
		x, y = y, x % y
	}
	return x
}
