package odd_int

import "fmt"

func FindOdd(seq []int) int {

	hashMap := make(map[int]int)

	for _, v := range seq {
		hashMap[v]++
	}

	fmt.Println("this is the seq: ", seq)
	fmt.Println("this is the hashmap: ", hashMap)
	for i, v := range hashMap {
		if isOdd(v) {
			return i
		}
	}
	return 0
}

func isOdd(x int) bool {

	if x%2 != 0 {
		return true
	}
	return false
}

func FindOdd2(seq []int) (res int) {
	for _, n := range seq {
		res ^= n
	}
	return
}
