package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var b = []byte("123")
	//string_internals(b)
	x := 1

	y := &x



	fmt.Println(*y)



	*y = 2

	fmt.Println(x)

	const (

		i = 7

		j

		k

	)
	fmt.Println(i, j, k)



	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}



	// creates new slice

	s := a[2:4]

	fmt.Println(s)



	//var v []int

	v := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = v[2:4]

	fmt.Println(cap(s))
	fmt.Println(len(s))

	fmt.Println(&v[2] == &s[0])

	s[0] = 33

	fmt.Println(a[2])

	newS := append(s, 55, 66)

	fmt.Printf("len=%d, cap=%d\n", len(newS), cap(newS))


	h := []int{1, 2, 3, 4, 5, 6}

	h = append(h, 7, 8)

	fmt.Println(len(h))

	//fmt.Printf("len=%d, len(h))


	s2 := []int{1, 2, 3}

	s3 := []int{4, 5, 6, 7}

	n2 := copy(s2, s3)

	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)



}

//to convert b from []byte to string without causing memory copy
func string_internals(b []byte) {
	s := *(*string)(unsafe.Pointer(&b))

	b[1] = '4'
	fmt.Printf("%+v. And the type is: %T\n ", s, s)
}


