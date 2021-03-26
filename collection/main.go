package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	A := []int{1, 2, 4, 5, 9, 2, 9, 1}
	fmt.Println("The sort: ", array_sort(A))
	fmt.Println("The unique: ", remove_duplicate(A))
	fmt.Println("highest: ", higest_num(A))
	fmt.Println("lowest: ", lowest_num(A))
	fmt.Println("lowest: ", reverse(A))

	B := []int{2, 3, 4, 6, 8, 10}
	fmt.Println("lowest: ", FindOutlier(B))

	starsTriangleForward(5)
	NumPats(5)
	NumNotAssigned(5)
	alphabet(5)
	alphabetCont(5)
	startAgain(5)
	fmt.Println("The factorial: ", fac(5))
	fmt.Println("The fibonacci: ", fib(6))
	fmt.Println("The prime_factors: ", prime_factors(6))

	fmt.Println("shortest word: ", shortestWord("normal_http_call my a name isdfs assdf"))

	fmt.Println(string_ending("abc", "c"))

	fmt.Println("This is the lcs: ", LCS("AGGTAB", "GXTXAYB"))

	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	fmt.Println(Fibonacci([2]float64{1, 1}, 10))
	fmt.Println(print_error("aaaxbbbbyyhwawiwjjjwwm"))

	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(bookseller(b, c))

	fmt.Println(decodeMorse(".... . -.--   .--- ..- -.. ."))

	fmt.Println(morse(".... . -.--   .--- ..- -.. ."))

}

func array_sort(arr []int) []int {

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func remove_duplicate(arr []int) []int {

	keys := make(map[int]bool)
	newArray := []int{}

	for _, v := range arr {
		if _, value := keys[v]; !value {
			keys[v] = true
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func higest_num(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func lowest_num(arr []int) int {
	low := arr[0]
	for _, v := range arr {
		if v < low {
			low = v
		}
	}
	return low
}

func reverse(arr []int) []int {

	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1

		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func FindOutlier(arr []int) int {

	odd_arr := []int{}
	even_arr := []int{}
	outlier := 0

	for _, v := range arr {
		if v%2 == 0 {
			even_arr = append(even_arr, v)
		} else {
			odd_arr = append(odd_arr, v)
		}
	}
	if len(odd_arr) > len(even_arr) {
		outlier = even_arr[0]
	} else {
		outlier = odd_arr[0]
	}

	return outlier
}

// PYRAMIDS

func starsTriangleForward(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func NumPats(n int) {
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", num)
		}
		num++
		fmt.Println()
	}
}

func NumNotAssigned(n int) {
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
}

// Character Pattern
func alphabet(n int) {
	rune := 'A'
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", string(rune))
		}
		rune++
		fmt.Println()
	}
}
func alphabetCont(n int) {
	rune := 'A'
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", string(rune))
			rune++
		}
		fmt.Println()
	}
}

func startAgain(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func fac(n int) int {
	if n < 2 {
		return 1
	}
	return n * fac(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func prime_factors(n int) []int {
	divisor := 2
	var factors []int
	for n > 1 {
		if n%divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		} else {
			divisor++
		}
	}
	return factors
}

func prime_factors2(n int) []int {
	divisor := 2
	var factors []int
	for n > 1 {
		if n%divisor == 0 {
			factors = append(factors, divisor)
			n = n % divisor
		} else {
			divisor++
		}
	}
	return factors
}

//shortest word:
func shortestWord(s string) int {
	if len(s) <= 0 {
		return 0
	}
	//convert the sentence to an array:
	words := strings.Split(s, " ")
	shortest := words[0]
	for _, v := range words {
		if len(v) < len(shortest) {
			shortest = v
		}
	}
	return len(shortest)
}

//String Ending
func string_ending(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	//   return strings.HasSuffix(str, ending)
	return str[len(str)-len(ending):] == ending
}

//Subsequence
func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)

	if m == 0 || n == 0 {
		return ""
	} else if s1[m-1] == s2[n-1] {
		return LCS(s1[:m-1], s2[:n-1]) + string(s1[m-1])
	}
	v1 := LCS(s1[:m-1], s2)
	v2 := LCS(s1, s2[:n-1])

	if len(v1) > len(v2) {
		return v1
	}
	return v2
}

func Tribonacci(signature [3]float64, n int) []float64 {
	result := signature[:]
	if n <= 3 {
		return result[:n]
	}
	for i := 3; i < n; i++ {
		sum := result[i-3] + result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

func Fibonacci(signature [2]float64, n int) []float64 {
	result := signature[:]
	if n <= 2 {
		return result[:n]
	}
	for i := 2; i < n; i++ {
		sum := result[i-2] + result[i-1]
		result = append(result, sum)
	}
	return result
}

func print_error(s string) string {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'm' {
			continue
		} else {
			count++
		}
	}
	return fmt.Sprintf("%d/%d", count, len(s))
}

func bookseller(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	result := ""
	for _, v := range listCat {
		tot := 0
		for _, w := range listArt {
			if w[0] == v[0] {
				ans, _ := strconv.Atoi(strings.Split(w, " ")[1])
				tot += ans
			}
		}
		if result != "" {
			result += " - "
		}
		result += "(" + v + " : " + strconv.Itoa(tot) + ")"
	}
	return result
}

func jose(n, k int) int {
	survivor := 0
	for i := 1; i < n; i++ {
		survivor = (survivor + k) % (i + 1)
	}
	return survivor + 1
}

var MORSE_CODE = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

func decodeMorse(morseCode string) string {
	decoded := ""
	morseWords := strings.Split(morseCode, "   ")
	for _, v := range morseWords {
		morseChars := strings.Split(v, " ")
		for _, char := range morseChars {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}

func morse(morseCode string) string {
	decoded := ""
	morseWords := strings.Split(morseCode, "   ")
	for _, v := range morseWords {
		morseChars := strings.Split(v, " ")
		for _, char := range morseChars {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}
