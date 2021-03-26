package main

import (
	"fmt"
	"strconv"
	"strings"
)

//You have a positive number n consisting of digits. You can do at most one operation: Choosing the index of a digit in the number, remove this digit at that index and insert it back to another or at the same place in the number in order to find the smallest number you can get.
//
//#Task: Return an array or a tuple or a string depending on the language (see "Sample Tests") with
//
//1) the smallest number you got
//2) the index i of the digit d you took, i as small as possible
//3) the index j (as small as possible) where you insert this digit d to have the smallest number.
//Example:
//
//smallest(261235) --> [126235, 2, 0] or (126235, 2, 0) or "126235, 2, 0"
//126235 is the smallest number gotten by taking 1 at index 2 and putting it at index 0
//
//smallest(209917) --> [29917, 0, 1] or ...
//
//[29917, 1, 0] could be a solution too but index `i` in [29917, 1, 0] is greater than
//index `i` in [29917, 0, 1].
//29917 is the smallest number gotten by taking 2 at index 0 and putting it at index 1 which gave 029917 which is the number 29917.
//
//smallest(1000000) --> [1, 0, 6] or ...

func main() {
	fmt.Println(Smallest(4232132))

	fmt.Println(Small(4232132))

}

func Small(n int64) []int64 {
	numStr := strconv.FormatInt(n, 10)
	//fmt.Printf("the numStr: %v and the type: %T\n", numStr, numStr) //is a string
	var ans int64 = n
	var ansI int64 = 0
	var ansJ int64 = 0
	strLen := len(numStr)
	for i := 0; i < strLen; i++ {
		for j := 0; j < strLen; j++ {
			arr := strings.Split(numStr, "") //covert the string to an array
			//fmt.Printf("the arr: %v and the type: %T\n", arr, arr) //it is a slice of string
			ch := arr[i]
			//fmt.Println("the c: ", ch)
			arr[i] = ""
			if i > j {
				arr[j] = ch + arr[j]
			} else {
				arr[j] = arr[j] + ch
			}
			newStr := strings.Join(arr, "")
			newInt, _ := strconv.ParseInt(newStr, 10, 64)
			if newInt < ans {
				ans = newInt
				ansI = int64(i)
				ansJ = int64(j)
			}
		}
	}
	return []int64{ans, ansI, ansJ}
}

func Smallest(n int64) []int64 {
	numStr := strconv.FormatInt(n, 10)
	var ans int64 = n
	var ansI int64 = 0
	var ansJ int64 = 0
	for i := 0; i < len(numStr); i++ {
		for j := 0; j < len(numStr); j++ {
			arr := strings.Split(numStr, "")
			ch := arr[i]
			arr[i] = ""
			if i > j {
				arr[j] = ch + arr[j]
			} else {
				arr[j] = arr[j] + ch
			}
			newStr := strings.Join(arr, "")
			newInt, _ := strconv.ParseInt(newStr, 10, 64)
			if newInt < ans {
				ans = newInt
				ansI = int64(i)
				ansJ = int64(j)
			}
		}
	}
	return []int64{ans, ansI, ansJ}
}
