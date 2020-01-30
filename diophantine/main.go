package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("this is the output: ", Solequa(90005))
}

func Solequa(n int) [][]int {
	r := make([][]int, 0)
	for i := 1; i < int(math.Sqrt(float64(n))) + 1; i++ {
		if n % i == 0 {
			d := n/i
			if (d - i) % 4 == 0 {
				y := (d - i) / 4
				x := i + 2 * y
				r = append(r, []int{x,y})
			}
		}
	}
	return r
}

func Solequa2(n int) (ret [][]int) {
	ret = [][]int{}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n % i != 0 {
			continue
		}
		t := n / i
		t1 := t - i
		if t1 % 4 != 0 {
			continue
		}
		ret = append(ret, []int{(t + i) / 2, t1 / 4})
	}
	return
}