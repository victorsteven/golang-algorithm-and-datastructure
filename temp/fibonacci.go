package temp

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	return result
}

func loopFib(n int) int {
	fibBox := []int{0,1}
	for i := 0; i < n; i++ {
		value := fibBox[i] + fibBox[i+1]
		fibBox = append(fibBox, value)
	}
	return fibBox[n]
}
