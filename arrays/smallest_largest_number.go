package arrays

func smallest_number(n []int) int {
	smallest := n[0]
	for _, v := range n {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func largest_number(n []int) int {
	largest := n[0]
	for _, v := range n {
		if v > largest {
			largest = v
		}
	}
	return largest
}
