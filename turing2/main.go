//package main
//
//
//
//import "fmt"
//
//
//
//func defFooStart() {
//
//	fmt.Println("defFooStart() executed")
//
//}
//
//
//
//func defFooEnd() {
//
//	fmt.Println("defFooEnd() executed")
//
//}
//
//
//
//func defMain() {
//
//	fmt.Println("defMain() executed")
//
//}
//
//
//
//func foo() {
//
//	fmt.Println("foo() executed")
//
//
//	defer defFooStart() // defer defFooStart call
//
//
//	panic("panic from foo()")
//
//
//	defer defFooEnd() // defer defFooEnd call
//
//
//	fmt.Println("foo() done")
//
//}
//
//
//
//func main() {
//
//	fmt.Println("main() started")
//
//
//	defer defMain() // defer defMain call
//
//
//	foo() // call foo function
//
//
//	fmt.Println("main() done")
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func greet(c chan string) {
//
//	fmt.Println("Hello " + <-c + "!")
//
//}
//
//
//
//func main() {
//
//	fmt.Println("main() started")
//
//	c := make(chan string)
//
//
//
//	go greet(c)
//
//
//
//	c <- "John"
//
//	fmt.Println("main() stopped")
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	fmt.Println("A")
//
//
//
//	c := make(chan string)
//
//	c <- "John"
//
//
//	fmt.Println("B")
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	c := make(chan int, 3)
//
//	c <- 1
//
//	c <- 2
//
//
//
//	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
//
//}

//package main
//
//
//
//import (
//
//	"fmt"
//
//	"runtime"
//
//)
//
//
//
//func squares(c chan int) {
//
//	for i := 0; i < 4; i++ {
//
//		num := <-c
//
//		fmt.Println(num * num)
//
//	}
//
//}
//
//
//
//func main() {
//
//	c := make(chan int, 3)
//
//	go squares(c)
//
//
//
//	fmt.Print(runtime.NumGoroutine())
//
//	c <- 1
//
//	c <- 2
//
//	c <- 3
//
//	c <- 4
//
//
//
//	fmt.Print( runtime.NumGoroutine())
//
//}

//package main
//
//
//
//import (
//
//	"fmt"
//
//	"time"
//
//)
//
//
//
//var start time.Time
//
//
//
//func init() {
//
//	start = time.Now()
//
//}
//
//
//
//func service1(c chan string) {
//
//	c <- "Hello from service 1"
//
//}
//
//
//
//func service2(c chan string) {
//
//	c <- "Hello from service 2"
//
//}
//
//
//
//func main() {
//
//	chan1 := make(chan string)
//
//	chan2 := make(chan string)
//
//
//
//	go service1(chan1)
//
//	go service2(chan2)
//
//
//
//	select {
//
//	case res := <-chan1:
//
//		fmt.Println("Response from service 1")
//
//		fmt.Println("mine: ", res)
//
//
//	case res := <-chan2:
//
//		fmt.Println("Response from service 2")
//
//		fmt.Println("mine: ", res)
//
//	default:
//
//		fmt.Println("No response received")
//
//	}
//
//}

//package main
//
//
//
//type S struct{}
//
//
//
//func (s S) F() {}
//
//
//
//type IF interface {
//
//	F()
//
//}
//
//
//
//func InitType() S {
//
//	var s S
//
//	return s
//
//}
//
//
//
//func InitPointer() *S {
//
//	var s *S
//
//	return s
//
//}
//
//func InitEfaceType() interface{} {
//
//	var s S
//
//	return s
//
//}
//
//
//
//func InitEfacePointer() interface{} {
//
//	var s *S
//
//	return s
//
//}
//
//
//
//func InitIfaceType() IF {
//
//	var s S
//
//	return s
//
//}
//
//
//
//func InitIfacePointer() IF {
//
//	var s *S
//
//	return s
//
//}
//
//
//
//func main() {
//
//	//println(InitType() == nil)
//
//	print(InitPointer() == nil)
//
//	print(InitEfaceType() == nil)
//
//	print(InitEfacePointer() == nil)
//
//	print(InitIfaceType() == nil)
//
//	print(InitIfacePointer() == nil)
//
//}
//
//
//
//package main
//
//func main() {
//	m := make(map[string]int)
//	m["a"] = 1
//	if v, ok := m["b"]; ok {
//		println(v)
//	}
//}

//package main
//
//
//
//const N = 3
//
//
//
//func main() {
//
//	m := make(map[int]*int)
//
//
//
//	for i := 0; i < N; i++ {
//
//		m[i] = &i
//
//	}
//
//
//
//	for _, v := range m {
//
//		print(*v)
//
//	}
//
//}

//package main
//
//
//
//const N = 3
//
//
//
//func main() {
//
//	m := make(map[int]*int)
//
//
//
//	for i := 0; i < N; i++ {
//
//		j := int(i)
//
//		m[i] = &j
//
//	}
//
//
//
//	for _, v := range m {
//
//		print(*v)
//
//	}
//
//}

//package main
//
//
//
//import "log"
//
//
//
//func f() {
//
//	defer func() {
//
//		if r := recover(); r != nil {
//
//			log.Printf("recover:%#v", r)
//
//		}
//
//	}()
//
//	panic(1)
//
//	panic(2)
//
//}
//
//
//
//func main() {
//
//	f()
//
//}

//package main
//
//
//
//const N = 10

//func main() {
//
//	m := make(map[int]int)
//
//
//
//	wg := &sync.WaitGroup{}
//
//	mu := &sync.Mutex{}
//
//	wg.Add(N)
//
//	for i := 0; i < N; i++ {
//
//		go func() {
//
//			defer wg.Done()
//
//			mu.Lock()
//
//			m[i] = i
//
//			mu.Unlock()
//
//		}()
//
//	}
//
//	wg.Wait()
//
//	println(len(m))
//
//}

//package main
//
//func main() {
//
//s := "123"
//
//ps := &s
//
//b := []byte(*ps)
//
//pb := &b
//
//
//
//s += "4"
//
//*ps += "5"
//
//b[1] = '0'
//
//
//
//print(*ps)
//
//print(string(*pb))
//
//}

//package main
//
//
//
//import (
//
//"math/rand"
//
//"sync"
//
//)
//
//
//
//const N = 10
//
//
//
//func main() {
//
//	m := make(map[int]int)
//
//
//
//	wg := &sync.WaitGroup{}
//
//	wg.Add(N)
//
//	for i := 0; i < N; i++ {
//
//		go func() {
//
//			defer wg.Done()
//
//			m[rand.Int()] = rand.Int()
//
//		}()
//
//	}
//
//	wg.Wait()
//
//	println(len(m))
//
//}
//
//
//
//Which of the following is the correct way to fix this issue so that the program avoids the "concurrent map writes" error?

//package main
//
//
//
//type S struct {
//
//	name string
//
//}

//
//func main() {
//
//	m := map[string]S{"x": S{"one"}}
//
//	m["x"].name = "two"
//
//	fmt.println(m[“x”))
//
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	var a0 []int = make([]int, 0)
//
//	fmt.Println(a0 == nil)
//
//	var ag []int = nil
//
//	fmt.Println(ag == nil)
//
//
//a := []string{"A", "B", "C", "D", "E"}
//
////a = nil
////
////fmt.Println(a, len(a), cap(a))
//
//
//	a = a[:0]
//
//	fmt.Println(a, len(a), cap(a))
//
//
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	for {
//
//		fmt.Println("a")
//
//	}
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	var s *int
//
//	fmt.Println(s)
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	var y = 458
//
//	var p = &y
//
//	fmt.Println(*p)
//
//}
//
//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	var y = 458
//
//	var p = &y
//
//	*p = 500
//
//	fmt.Println(y)
//
//}

//package main
//
//
//
//import "fmt"
//
//
//
//func main() {
//
//	val1 := 2345
//
//	val2 := 567
//
//
//
//	var p1 *int
//
//	p1 = &val1
//
//	p2 := &val2
//
//	res1 := &p1 == &p2
//
//	fmt.Println(res1)
//
//}

//package main
//
//
//
//import "fmt"
//
//func main() {
//
//	mychnl := make(chan string, 5)
//
//	mychnl <- "GFG"
//
//	mychnl <- "gfg"
//
//	mychnl <- "Turing"
//
//	mychnl <- "Developers"
//
//	fmt.Println( cap(mychnl))
//
//}

package main

import "fmt"

func main() {

	pair1 := [2]int{4, 2}

	pair2 := [2]int{4, 2}

	if pair1 != pair2 {

		fmt.Println("different pair")

	} else {

		fmt.Println("same pair")

	}

}
