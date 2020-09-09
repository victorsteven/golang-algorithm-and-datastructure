package rotate_for_max

import (
	"fmt"
	"strconv"
)

//Let us begin with an example:
//
//Take a number: 56789. Rotate left, you get 67895.
//
//Keep the first digit in place and rotate left the other digits: 68957.
//
//Keep the first two digits in place and rotate the other ones: 68579.
//
//Keep the first three digits and rotate left the rest: 68597. Now it is over since keeping the first four it remains only one digit which rotated is itself.
//
//You have the following sequence of numbers:
//
//56789 -> 67895 -> 68957 -> 68579 -> 68597
//
//and you must return the greatest: 68957.
//
//Task
//Write function max_rot(n) which given a positive integer n returns the maximum number you got doing rotations similar to the above example.
//
//So max_rot (or maxRot or ... depending on the language) is such as:
//
//max_rot(56789) should return 68957
//
//max_rot(38458215) should return 85821534

func MaxRot(n int64) int64 {

	fmt.Println("given: ", n)

	s, max := strconv.FormatInt(n, 10), n
	for i := 0; i < len(s); i++ {
		//s = s[:i] + s[i:][1:] + s[i:][:1]
		s = s[:i] + s[i+1:] + s[i:][:1]
		v, _ := strconv.ParseInt(s, 10, 64)
		if max < v {
			max = v
		}
	}
	return max
}

func MaxRot2(n int64) (max int64) {

	max = n
	strn := strconv.FormatInt(n, 10)

	var left string
	var middle string
	var right string
	var currentNumber int64

	fmt.Println("the iterating string: ", strn[:len(strn)-1])

	for index := range strn[:len(strn)-1] {

		left = strn[:index]
		middle = strn[index:index+1]
		right = strn[index+1:]

		fmt.Println("initial left, middle and right: ", left, middle, right)

		strn = left + right + middle

		currentNumber, _ = strconv.ParseInt(strn, 10, 64)

		if max < currentNumber {
			max = currentNumber
		}
	}
	return
}


//func MaxRot(n int64) int64 {
//	e := fmt.Sprint(n)
//	arr := []int{}
//	arr = append(arr, int(n))
//	for i := 0; i < len(e)-1; i++ {
//		e = rotate(e, i)
//		i, _ := strconv.Atoi(e)
//		arr = append(arr, i)
//	}
//	res := arr[0]
//	for _, e := range arr {
//		if e >= res {
//			res = e
//		}
//	}
//	return int64(res)
//}
//
//func rotate(n string, i int) string {
//	return n[:i] + n[i+1:] + string(n[i])
//}