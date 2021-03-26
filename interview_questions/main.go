package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		//fmt.Println("a value: ", v)
		v += 10
		//fmt.Println("a value update: ", v)
	}
	fmt.Println("Final: ", ss) //this will still give the previous values

	for i := range ss {
		//fmt.Println("a value before: ", ss[i])
		ss[i] += 10
		//fmt.Println("a value update: ", ss[i])
	}
	fmt.Println("FInal 2: ", ss) //this will be previous value plus 10, because we updated the previous value

	m := make(map[string]int) //A
	m["a"] = 1
	if v, ok := m["a"]; ok { //B
		println(v)
	}

	p := *f()
	print(p.m)

	println()
	printCorrect()

	println()
	outerForloop()

	println()
	goRoutineClosure()

	sa := "123"
	ps := &sa
	fmt.Printf("The type: %T\n", ps)

	println()
	interfaceEqual()

	println()
	mapImmutability()

	println()
	sortStruct()
}

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}
}

//print out the right value
func printCorrect() {
	const N = 3
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		j := i
		m[i] = &j
	}
	for _, v := range m {
		print(*v)
	}
}

//Exit the outer forloop
func outerForloop() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, "")
			break outer
		}
		println()
	}
}

func goRoutineClosure() {

	const N = 10
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}

func interfaceEqual() {

	type S struct {
		a, b, c string
	}
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(reflect.DeepEqual(x, y))
}

func mapImmutability() {
	type S struct {
		name string
	}
	m := map[string]*S{"x": &S{"one"}}
	m["x"].name = "two"

	//fmt.Println("the m: ", &m)
}

func sortStruct() {
	type S struct {
		v int
	}
	s := []S{{1}, {3}, {5}, {2}}
	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})
	fmt.Printf("%#v", s)
}
