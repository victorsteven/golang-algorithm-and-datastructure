package caeser_cipher_test

import (
	"efficient-algorithms/caeser_cipher"
	"testing"
)

func assert(expected, actual string) bool {

	if expected == actual {
		return true
	}
	return false
}

//caesar("Abcd", 2) should return "Cdef"
//caesar("message", -1) should return "ldrrzfd"
//caesar("ZZ Top", 3) should return "CC Wrs"

func TestCipher(t *testing.T) {

	result1 := caeser_cipher.Cipher("Abcd", 2)
	result2 := caeser_cipher.Cipher("message", -1)
	result3 := caeser_cipher.Cipher("ZZ Top", 3)

	if ok := assert("Cdef", result1); !ok {
		t.Errorf("got %s, want %s", result1, "Cdef")
	}

	if ok := assert("ldrrzfd", result2); !ok {
		t.Errorf("got %s, want %s", result2, "ldrrzfd")
	}

	if ok := assert("CC Wrs", result3); !ok {
		t.Errorf("got %s, want %s", result3, "CC Wrs")
	}
}
