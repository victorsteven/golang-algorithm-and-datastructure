package main

import "fmt"

func main() {
	fmt.Println(grider(10))
}

func grider(rows int) [][]int {
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, 4) //each row have 4 columns
	}
	return grid
}
