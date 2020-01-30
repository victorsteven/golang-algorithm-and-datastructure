package main

import "fmt"

func main() {
	var s = make([]int, 3)
	n := copy(s, []int{0,1,2,3})


	fmt.Println(n)
	fmt.Println("the slice: ", s)

	b := make([]byte, 5)
	m := copy(b, "Hello world")
	fmt.Println(m)
	fmt.Println("the byte slice: ", string(b))


	var s1 []int
	s2 := []int{1, 2, 3}
	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil", s1 == nil)
}
