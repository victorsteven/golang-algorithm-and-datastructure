package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("this is the print: ", numPlayers(8, []int32{50, 51, 23, 25, 37, 65, 55, 5, 98, 93}))
 fmt.Println("-------------")
	fmt.Println("this is the print 2: ", numPlayers(4, []int32{2,2,2,4,5}))

	fmt.Println("this is the print 3: ", numPlayers(4, []int32{2,2,3,4,5}))

	fmt.Println("this is the print 4: ", numPlayers(4, []int32{2,2,2,4,5}))
}


func numPlayers(k int32, scores []int32) int32 {
	// Write your code here
	size := len(scores)

	divisor := int32(size) * k

	var count int32

	// if the divisor is less than the least member, dont use it
	if isDivisible(divisor, scores) {
		count = divideAndCount(divisor, k, scores)
	//} else {
		count = onlyCount(k, scores)
	}

	return count

}

func isDivisible(div int32, scores []int32) bool {

	for _, v := range scores {
		if div > v {
			return false
		}
	}

	return true
}


func divideAndCount(divisor, k int32, scores []int32) int32 {

	fmt.Println("this is run divide")
	var count int32
	// loop through the array and divide by the divisor:
	for _, v := range scores {
		if v/divisor <= k {
			count++
		}
	}
	return count
}

func onlyCount(k int32, scores []int32) int32 {

	fmt.Println("this is run count")
	var count int32
	// loop through the array and divide by the divisor:

	// sort the array:
	scoresInts := scoresInts(scores)

	sort.Ints(scoresInts)

	result := []int{}

	fmt.Println("the sorted man: ", scoresInts)
	fmt.Println("this is the k: ", k)

	resultMap := make(map[int]bool)

	for i, v := range scoresInts {
		if int(k) >= i+1   {
			fmt.Println("our value: ", i)
			//count++
			result = append(result, v)

		}
	}

	fmt.Println("the result arr: ", result)

	for i, _ := range result {
		if int(k) >= i   {
			count++
		}
	}
	////for i := 0; i <= len(scoresInts); i++ {
	//	//fmt.Println("position: ", i)
	//	//index := i + 1
	//	if int(k) >= i  {
	//		fmt.Println("the element and index: ", scoresInts[i], i)
	//		//count++
	//		//result :=
	//		result = append(result, i + 1)
	//	}
	//}

	return count
}


func scoresInts(scores []int32) (scoresInt []int) {
	for _, v := range scores {
		scoresInt = append(scoresInt, int(v))
	}
	return scoresInt
}