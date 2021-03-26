package main

import (
	"fmt"
	"sync"
)

func makeSquares(array [10]int) {
	for index, elem := range array {
		array[index] = elem * elem
	}
}

func makeSquaresInt(slice []int) {
	for index, elem := range slice {
		slice[index] = elem * elem
	}
}

func main() {
	s10 := make([]int, 2, 4)
	fmt.Printf("The print: s10=%v\n", s10)
	sy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sy[2:])
	fmt.Println(sy[:4])
	fmt.Println(sy[2:4])
	fmt.Println(sy[:])
	sc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	makeSquaresInt(sc)
	fmt.Println(sc)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	makeSquares(a)
	fmt.Println("the a's: ", a) //the remain the same as the above a, i think because an array was used in its square method

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	//s1 := []int{0, 1, 2, 3}
	//s2 := []int{0, 1, 2, 3}
	//fmt.Println(s1 == s2) //compilation error. you dont check slices this way

	var sp Shape
	fmt.Println("value of s is", sp)
	fmt.Printf("type of s is %T\n", sp)

	var sa Shape
	sa = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Println(sa == r) //true

	var sb Shape = Rect{10, 3}
	fmt.Println(sb)

	fmt.Println("main() started")
	var wg sync.WaitGroup
	//wg.Add(1)
	go service(wg)

	wg.Wait()
	select {}
	fmt.Println("main() stopped")
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func service(wg sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Hello from service!")
	wg.Done()

}
