package main

import "fmt"

func main() {
	fmt.Println(intCast(23))
	fmt.Println(stringCast(23))
}

//checking if the interface input type is an int
func intCast(x interface{}) int {
	if v, ok := x.(int); ok {
		return v
	} else {
		return -1 //if not, return -1
	}
}

//checking if the interface input type is a string
func stringCast(x interface{}) string {
	if str, ok := x.(string); ok {
		return str
	} else {
		return "cant be converted"
	}
}
