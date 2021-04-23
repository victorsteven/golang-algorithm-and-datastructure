package palindrome_number

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {

	if isPalindrome(121) == false {
		t.Errorf("121 is not a palindrome")
	}
}

func TestPalindromeNumber2(t *testing.T) {

	if isPalindrome(-121) == true {
		t.Errorf("-121 is not a palindrome")
	}
}

func TestPalindromeNumber3(t *testing.T) {

	if isPalindrome(10) == true {
		t.Errorf("10 is not a palindrome")
	}
}

func TestPalindromeNumber4(t *testing.T) {

	if isPalindrome(-101) == true {
		t.Errorf("-101 is not a palindrome")
	}
}
