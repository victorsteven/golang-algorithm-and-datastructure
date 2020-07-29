package main

import (
	"fmt"
	"math"
	"sort"
)

//Create a function that takes a positive integer and returns the next bigger number that can be formed by rearranging its digits. For example:
//12 ==> 21
//513 ==> 531
//2017 ==> 2071


func main() {

	fmt.Println(NextBigger(3452))

}

func intToSlice(n int) []int {
	var numb []int
	for n > 0 {
		numb = append(numb, n % 10)
		n /= 10
	}
	return numb
}

func sliceToInt(numb []int) int {
	n := 0
	for i, v := range numb {
		//fmt.Println("the ints: ", int(math.Pow10(i))) //1, 10, 100, 1000
		n += v * int(math.Pow10(i))
	}
	return n
}

func NextBigger(n int) int {
	numb := intToSlice(n)
	i := 0
	for i := 0; i < len(numb); i++ {
		if numb[i] > numb[i + 1] {
			break
		}
	}
	if i == len(numb)-1 {
		return -1
	}

	mn := i
	for j := 0; j < i; j++ {
		if numb[j] > numb[i + 1] {
			mn = j
		}
	}
	numb[mn], numb[i + 1] = numb[i + 1], numb[mn]
	tmp := numb[:i + 1]
	//fmt.Println("the temp: ", tmp)
	sort.Slice(tmp, func(i,j int) bool {
		return tmp[i] > tmp[j]
	})
	return sliceToInt(numb)
}

//func NextBigger(n int) int {
//	numb := intToSlice(n)
//	i := 0
//	for i := 0; i < len(numb) - 1; i++ {
//		if numb[i] > numb[i + 1] {
//			break
//		}
//	}
//	if i == len(numb)-1 {
//		return -1
//	}
//
//	mn := i
//
//	for j := 0; j < i; j++ {
//		if numb[j] > numb[i + 1] {
//			mn = j
//		}
//	}
//	numb[mn], numb[i+1] = numb[i + 1], numb[mn]
//	tmp := numb[:i + 1]
//	sort.Slice(tmp, func(i, j int) bool {
//		return tmp[i] > tmp[j]
//	})
//	return sliceToInt(numb)
//}

func countDigit(i int) (count int) {
	for i != 0 {
		i = i /10
		fmt.Println(i)
		//i /= 10
		count++
	}
	return count
}
