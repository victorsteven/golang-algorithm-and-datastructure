package main

import "fmt"

func main() {
	ans := jose(7,3)
	fmt.Println(ans)
}

func josephus_survivor(people, elim int) int {
	if people == 1 {
		return 1
	} else {
		return (josephus_survivor(people - 1, elim) + elim - 1) % people + 1
	}
}

//Solution 2:
func jose(n, k int) int {
	survivor := 0
	for i := 1; i < n; i++ {
		survivor = (survivor + k) % (i + 1)
	}
	return survivor + 1
}




