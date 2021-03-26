package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))
	fmt.Println(PickPeaks([]int{1, 2, 3, 5, 2, 1, 5}))
	fmt.Println(PickPeaks([]int{1, 1, 2, 2, 2, 2, 1, 1}))
}

func PickPeaks(array []int) PosPeaks {
	posPeaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
	maxpos := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			maxpos = i
		} else if array[i-1] > array[i] && maxpos > 0 {
			posPeaks.Pos = append(posPeaks.Pos, maxpos)
			posPeaks.Peaks = append(posPeaks.Peaks, array[maxpos])
			maxpos = 0
		}
	}
	return posPeaks
}
