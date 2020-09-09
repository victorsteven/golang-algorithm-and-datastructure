package main

import (
	"fmt"
	"strings"
)

//This kata is the first of a sequence of four about "Squared Strings".
//
//You are given a string of n lines, each substring being n characters long: For example:
//
//s = "abcd\nefgh\nijkl\nmnop"
//
//We will study some transformations of this square of strings.
//
//Vertical mirror: vert_mirror (or vertMirror or vert-mirror)
//vert_mirror(s) => "dcba\nhgfe\nlkji\nponm"
//Horizontal mirror: hor_mirror (or horMirror or hor-mirror)
//hor_mirror(s) => "mnop\nijkl\nefgh\nabcd"
//or printed:
//
//vertical mirror   |horizontal mirror
//abcd --> dcba     |abcd --> mnop
//efgh     hgfe     |efgh     ijkl
//ijkl     lkji     |ijkl     efgh
//mnop     ponm     |mnop     abcd
//Task:
//Write these two functions
//and
//
//high-order function oper(fct, s) where
//
//fct is the function of one variable f to apply to the string s (fct will be one of vertMirror, horMirror)
//Examples:
//s = "abcd\nefgh\nijkl\nmnop"
//oper(vert_mirror, s) => "dcba\nhgfe\nlkji\nponm"
//oper(hor_mirror, s) => "mnop\nijkl\nefgh\nabcd"
//Note:
//The form of the parameter fct in oper changes according to the language. You can see each form according to the language in "Sample Tests".
//
//Bash Note:
//The input strings are separated by , instead of \n. The output strings should be separated by \r instead of \n. See "Sample Tests".
//
//Forthcoming katas will study other transformations.


func main() {
	str := "abcdef"
	fmt.Println("the original: ", str)
	arr := strings.Split(str, "")
	ans := reverse(arr)
	fmt.Println(ans)
	fmt.Printf("the type: %T\n", ans)
}

func reverse(s []string) string {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}

	return strings.Join(s, "")

}

func VertMirror(s string) string {

	var b strings.Builder
	b.Grow(len(s) + 1)

	n := strings.Index(s, "\n")

	for i := 0; i < len(s); i += n + 1 {
		for j := i + n - 1; j >= i; j-- {
			b.WriteByte(s[j])
		}
		b.WriteByte('\n')
	}
	v := b.String()
	v = v[:b.Len()-1]

	return v
}

func HorMirror(s string) string {

	var b strings.Builder
	b.Grow(len(s) + 1)

	n := strings.Index(s, "\n")

	for i := len(s) - n; i >= 0; i -= (n + 1) {
		b.WriteString(s[i : i+n])
		b.WriteByte('\n')
	}
	v := b.String()
	v = v[:b.Len()-1]

	return v
}

type FParam func(string) string
func Oper(f FParam, x string) string {
	return f(x)
}


///////////////////////////////////////////
func revstring(s string) string {
	var u = make([]byte, len(s))
	copy(u, s)
	for i, j := 0, len(u)-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return string(u)
}
func revarray(s [] string) [] string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func VertMirror2(s string) string {
	arr := strings.Split(s, "\n")
	var res = []string{}
	for i := 0; i < len(arr); i++ {
		res = append(res, revstring(arr[i]))
	}
	return strings.Join(res, "\n")
}
func HorMirror2(s string) string {
	arr := strings.Split(s, "\n")
	revarray(arr)
	return strings.Join(arr, "\n")
}


////////////////////////////////////////////////
func VertMirror3(s string) string {
	arr := strings.Split(s, "\n")
	f := ""
	for _, v := range arr {
		rs := []rune(v)
		for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
			rs[i], rs[j] = rs[j], rs[i]
		}
		f += string(rs) + "\n"
	}
	return fmt.Sprintf("%s", strings.TrimRight(f, "\n"))
}

func HorMirror3(s string) string {
	arr := strings.Split(s, "\n")
	f := ""
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for _, v := range arr {
		f += v + "\n"
	}
	return fmt.Sprintf("%s", strings.TrimRight(f, "\n"))
}