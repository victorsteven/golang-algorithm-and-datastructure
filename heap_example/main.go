package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//str := "hello, Sam, how are you today. Alice, good to me you"
	//fmt.Println(vowel(str))
	h := &intHeap{6,2,1,5}
	heap.Init(h)
	for h.Len() > 0 {
		fmt.Printf("%d", heap.Pop(h))
	}
}


type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}