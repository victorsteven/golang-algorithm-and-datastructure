package example_1


//Find all numbers less than n, which are palindromic. Numbers can be printed in any order.


//func createPalindrome(input, b int, isOdd bool) int {
//
//	n := input
//	palin := input
//
//	if isOdd {
//		n = n / b
//	}
//
//	//creates  palindrome by just appending reverse of number of itself
//	for n > 0 {
//		palin = palin * b + (n % b)
//		n = n / b
//	}
//	return palin
//}
//
//func generatePalindrome(n int) int {
//
//	num := 0
//
//	for j := 0; j < 2; j++ {
//
//		// Creates palindrome numbers
//		// with first half as i. Value
//		// of j decided whether we need
//		// an odd length of even length
//		// palindrome.
//		//i := 1
//		//for num = createPalindrome(i, 10, j % 2)) < n {
//		//	fmt.Println(n, ""),
//		//	i++
//		//}
//	}
//}