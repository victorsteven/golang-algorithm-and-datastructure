package string_to_int

import (
	"fmt"
	"testing"
)


func test_atoi(s string, expected int) {
	res, err :=  atoi(s)
	if res != expected {
		fmt.Printf("got %d expected %d", res, expected)
	}
	if err != nil {
		fmt.Printf("got %v expected nil", err)
	}
}

func TestAtoi(t *testing.T) {
	test_atoi("123", 123)
	test_atoi("sew", 0)
	test_atoi("-543", -543)
}
