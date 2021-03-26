package main

import "fmt"

//1. What is the result of re-slicing?
// A new reference and boundary over the original data

//2. What is true about an empty interface?
// A function can take an empty interface as a parameter

// 3. You wrote a comprehensive test suite for a library, including regular and benchmark tests, but the benchmark tests aren't executing. Why might that be?
// The -bench flag was not added to the go test command.

// this will print true:
//matched, _ := regexp.MatchString("food", "seafood")
//fmt.Println(matched)

// 13. Two projects, A and B, exist in your workspace and both depend on a third-party library. Because of the requirements of project A, the library is updated to the latest version. Unfortunately, the new version contains changes that break project B. From the given actions, what is the best way to address this, while having the least possible affect on the projects?
// Move the library to project-specific vendor libraries.

//func main() {
//	fmt.Println("this is the test platform")
//}

type keet struct {
	F string `species:"gopher" color:"blue"`
	G string `species:"hopper" color:"green"`
}

//func main() {
//
// a := 2
//
// af := float64(a)
//	k := keet{}
//	val := reflect.TypeOf(k)
//
//	next := val.Field(1)
//	fmt.Println(next.Tag.Get("color"), next.Tag.Get("species"))
//
//}

//func display(i int, e error) {
//	fmt.Printf("(%v, %v)\n", i, e)
//}
//func getValue(r int) (int, error) {
//	return r, nil
//}
//func main() {
//	foo := 1000
//	display(foo, nil)
//	if true {
//		foo, err := getValue(2000)
//		display(foo, err)
//	}
//	display(foo, nil)
//	if true {
//		var err error
//		foo, err = getValue(3000)
//		display(foo, err)
//	}
//	display(foo, nil)
//}

func gotoUseCase() {
	for i := 0; i < 30; i++ {
	myNum:

		if i == 2 {

			goto myNum
		}
	}
}

//func main() {
//
//	//gotoUseCase()
//	//fmt.Println(87)
//
//	m, _ := regexp.MatchString("food", see)
//}

//func main() {
//	fmt.Println(len(os.Args))
//}

//type Dog struct {}
//
//
//func (d Dog) SpeakMore() {
//	fmt.Println("bark")
//}
//
//
//func (d *Dog) Speak() {
//	fmt.Println("woof")
//}
//
//
//func main() {
//	animal1 := new(Dog)
//	animal1.Speak()
//	animal1.SpeakMore()
//
//
//	var animal2 Dog
//	animal2.Speak()
//	animal2.SpeakMore()
//}

func main() {
	vals := make(chan float64, 4)

	vals <- 85.4
	vals <- 0.6
	vals <- -54.9
	vals <- 32.14

	if vals == vals {
		fmt.Println(3)
		if vals == vals {
			fmt.Println(0)
			if <-vals == <-vals {
				fmt.Println(5)
			} else {
				if <-vals == <-vals {
					fmt.Println(8)
				}
			}
		}
	}
}
