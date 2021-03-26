package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {

	//how to check if a map contains a key in go
	var mine = make(map[string]string)

	mine = map[string]string{
		"name": "Mike",
		"age":  "34",
	}

	fmt.Println("the map: ", mine)

	//check if a value is found in a map:
	if val, ok := mine["name"]; !ok {
		fmt.Println("No mike")
	} else {
		fmt.Println("this is the value: ", val)
	}

	//the map:
	myMap := []int{1, 3, 5, 3, 5, 6, 3, 5, 2, 1}
	fmt.Println(removeDup(myMap))

	a := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(reverse(a))

	sentence := "This is the man that shouted normal_http_call and the bro turned to him. Awesome bro how are you, the man asked. The boy replied normal_http_call"

	fmt.Println(countWords(sentence))

	str := "manner"
	fmt.Println(reverseWord(str))

	mapper := map[string]bool{"A": true, "B": true}
	fmt.Println(mapCopy(mapper))

	fmt.Println("this is the tag example below: ")
	metaFunc()

	fmt.Println()
	fmt.Println(checkEqual([]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}))

	fmt.Println()
	usingScanner()

	fmt.Println()
	fmt.Println(arrayJoin([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))

	fmt.Println()
	Inter()
}

//remove duplicates in an array

func removeDup(arr []int) []int {
	newArray := []int{}
	keys := make(map[int]bool)

	for _, v := range arr {
		if value := keys[v]; !value {
			//the value has not existed before
			keys[v] = true
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("value of i: ", i)
		fmt.Println("value of j: ", j)

		s[i], s[j] = s[j], s[i]
	}
	return s
}

func countWords(s string) int {
	count := 0
	w := strings.Split(s, " ")
	for _, v := range w {
		switch v {
		case "normal_http_call", "bro":
			count++
		}
	}
	return count
}

//word reversal:
func reverseWord(s string) string {
	w := strings.Split(s, "")
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	formedString := strings.Join(w, "")

	return formedString
}

//copy a map in Go
func mapCopy(a map[string]bool) map[string]bool {
	b := make(map[string]bool)
	for key, value := range a {
		b[key] = value
	}
	return b
}

func metaFunc() {
	type User struct {
		Name  string `mytag:"MyName"`
		Email string `mytag:"MyEmail"`
	}
	u := User{"Steven", "steven@gmail.com"}
	t := reflect.TypeOf(u)
	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag' : %q\n", field.Tag.Get("mytag"))

	}
}

//checking if two slices are equal

func checkEqual(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func usingScanner() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	//var (
	//	text string
	//	n    int
	//)

	//for scanner.Scan() {
	//	n++
	//	text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
	//}
	//fmt.Print(text)

	//A better way
	var (
		text []byte
		n    int
	)
	for scanner.Scan() {
		n++
		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
	}
	//fmt.Print(text)
	os.Stdout.Write(text)
}

func arrayJoin(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	return arr1
}

//comparing interfaces
func Inter() {
	var a interface{}
	var b interface{}

	a = []int{1}
	b = []int{2}
	fmt.Println(a == b) //this will lead to panic
}
