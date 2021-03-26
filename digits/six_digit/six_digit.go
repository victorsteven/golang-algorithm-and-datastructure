package main

//Given any six digit number, how could one output a slice where each character of that number is assigned an individual location within the slice?
//
//For instance, a slice (let's call it s) containing all of these characters, would have s[0]=first digit, s[1]=second digit, s[2]=third digit and so on.
//
//Any help would be greatly appreciated!

/////////////////
//func IntToSlice(n int64, sequence []int64) []int64 {
//	if n != 0 {
//		i := n % 10
//		sequence = append([]int64{i}, sequence...)
//
//		return IntToSlice(n/10, sequence)
//	}
//	return sequence
//}
//
//func main() {
//
//	ans := IntToSlice(3, []int64{1,2,3,4,5,6})
//
//	fmt.Println(ans)
//
//}
