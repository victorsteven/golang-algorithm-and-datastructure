package main

import (
	"fmt"
	"sort"
)

//Given a positive integer K, a circle center at (0, 0) and coordinates of some points. The task is to find minimum radius of the circle so that at-least k points lie inside the circle. Output the square of the minimum radius.

//Input : (1, 1), (-1, -1), (1, -1),
//         k = 3
//Output : 2
//We need a circle of radius at least 2
//to include 3 points.
//
//
//Input : (1, 1), (0, 1), (1, -1),
//         k = 2
//Output : 1
//We need a circle of radius at least 1
//to include 2 points. The circle around
//(0, 0) of radius 1 would include (1, 1)
//and (0, 1).

func minRadius(k int, x, y []int, n int) int {

	//this slice is not created with the length of 0, so we can add to it by "reslicing".
	//slices created with zero length can receive elements using the append() function, else, it will throw an error similar to this: "index out of range [0] with length 0"
	dis := make([]int, n)

	//Finding distance between of each point from origin
	for i := 0; i < n; i++ {

		dis[i] = x[i] * x[i] + y[i] * y[i]

		//sorting the distance
		sort.Ints(dis)

	}
	return  dis[k-1]
}

func main() {

	k := 3
	x := []int{1, -1, 1}
	y := []int{1, -1, -1}

	n := len(x)

	fmt.Println(minRadius(k, x, y, n))
}