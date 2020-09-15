package decimal_to_factorial

import (
	"strconv"
	"strings"
)

//Coding decimal numbers with factorials is a way of writing out numbers in a base system that depends on factorials, rather than powers of numbers.
//
//In this system, the last digit is always 0 and is in base 0!. The digit before that is either 0 or 1 and is in base 1!. The digit before that is either 0, 1, or 2 and is in base 2!, etc. More generally, the nth-to-last digit is always 0, 1, 2, ..., n and is in base n!.
//
//Read more about it at: http://en.wikipedia.org/wiki/Factorial_number_system
//
//Example
//The decimal number 463 is encoded as "341010", because:
//
//463 = 3×5! + 4×4! + 1×3! + 0×2! + 1×1! + 0×0!
//
//If we are limited to digits 0..9, the biggest number we can encode is 10!-1 (= 3628799). So we extend 0..9 with letters A..Z. With these 36 digits we can now encode numbers up to 36!-1 (= 3.72 × 1041)
//
//Task
//We will need two functions. The first one will receive a decimal number and return a string with the factorial representation.
//
//The second one will receive a string with a factorial representation and produce the decimal representation.
//
//Given numbers will always be positive.


func FacsString2Dec(str string) int {
	response := 0
	step := len(str)
	iter := 0

	for step > 0 {
		rest, _ := strconv.ParseInt(string(str[iter]), 16, 10)
		response = response * step + int(rest)
		step--
		iter++
	}
	return response
}

func Dec2FactString(nb int) string {

	var response string

	step := 1
	remainder := 0
	for nb > 0 {
		nb, remainder = nb/step, nb%step
		response = strconv.FormatInt(int64(remainder), 16) + response
		step++
	}
	return strings.ToUpper(response)
}


//func Dec2FactString(nb int) string {
//	var alphabet []string = strings.Split("0123456789ABCDEFGHIGKLMNOPQRSTUWXYZ", "")
//	var res = ""; base := 1
//	for nb > 0 {
//		res = alphabet[(nb % base)] + res
//		nb = int(math.Floor(float64(nb) / float64(base)))
//		base++
//	}
//	return res
//}
//func FactString2Dec(str string) int {
//	var alphabet string = "0123456789ABCDEFGHIGKLMNOPQRSTUWXYZ"
//	arr := strings.Split(str, ""); var res int = 0; var fact int = 1
//	for i := len(arr) - 2; i >= 0; i-- {
//		c := arr[i]
//		a := strings.Index(alphabet, c)
//		fact *= len(arr) - i - 1
//		res += a * fact
//	}
//	return res
//}