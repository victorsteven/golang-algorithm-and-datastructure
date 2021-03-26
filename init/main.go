// A/a1.go
package A

func init() {
	println("a1")
}

var A1 = ""

// A/a2.go
//package A
//func init() {
//	println("a2")
//}
//var A2 = ""
//
//
//// B/b1.go
//package B
//func init() {
//	println("b1")
//}
//var B1 = ""
//
//
//// B/b2.go
//package B
//import "github.com/test/A"
//func init() {
//	println("b2")
//}
//func f() {
//	_ = A.A2
//}
//var B2 = ""
//
//
//// C/main.go
//package main
//import (
//"github.com/test/B"
//)
//func main() {
//	_ = B.B2
//}
