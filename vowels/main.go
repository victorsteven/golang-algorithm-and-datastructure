package main

import "fmt"

func main() {
	fmt.Println(vowel("hellohellohelloa"))

	//fmt.Println("ans: ", vowel("I wonder how many vowels?"))

	//if vowel("I wonder how many vowels?") == 7 && vowel("Ethan likes cheddar cheese") == 9 {
	//	fmt.Println(":) Correct! :)")
	//} else {
	//	fmt.Println(":(")
	//}

}

func vowel(myString string) int {
	count := 0
	//for _, value := range myString {
	//	switch value {
	//	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
	//		count++
	//	}
	//}

	for _, v := range myString {
		switch {
		case string(v) == "a":
			count++
		case string(v) == "e":
			count++
		case string(v) == "i":
			count++
		case string(v) == "o":
			count++
		case string(v) == "u":
			count++
		case string(v) == "A":
			count++
		case string(v) == "E":
			count++
		case string(v) == "I":
			count++
		case string(v) == "O":
			count++
		case string(v) == "U":
			count++
		}
	}
	return count
}
