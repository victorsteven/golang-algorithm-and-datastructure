package phone_number

import (
	"fmt"
)

//Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

func CreatePhoneNumber(numbers [10]uint) string {

	for _, v := range numbers {
		if v > 9 || v < 0 {
			return ""
		}
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}
