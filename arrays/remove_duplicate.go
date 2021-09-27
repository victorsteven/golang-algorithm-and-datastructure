package arrays

func remove_duplicate(n []int) []int {
	var hashMap = make(map[int]bool)
	var newArray = make([]int, 0)
	for _, v := range n {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = true
			newArray = append(newArray, v)
		}
	}
	return newArray
}
