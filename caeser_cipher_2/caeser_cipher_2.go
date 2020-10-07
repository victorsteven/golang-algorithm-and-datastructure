package main

import (
	"fmt"
	"strings"
)

type ckey struct {
	enc, dec func(rune) rune
}

func newCaeser(k int) (*ckey, bool) {
	if k < 1 || k > 25 {
		return nil, false
	}

	rk := rune(k)

	return &ckey{
		enc: func(c rune) rune {
			if c >= 'a' && c < 'z' - rk || c >= 'A' && c <= 'Z' - rk {
				return c + rk
			} else if c > 'z' - rk && c <= 'z' || c > 'Z' - rk && c <= 'Z' {
				return c + rk - 26
			}
			return c
		},
		dec: func(c rune) rune {
			if c >= 'a' + rk && c <= 'z' || c >= 'A' + rk && c <= 'Z' {
				return c - rk
			} else if c >= 'a' && c < 'a' + rk || c >= 'A' && c < 'A' + rk {
				return c - rk + 26
			}
			return c
		},
	}, true
}


func (ck ckey) encipher(pt string) string {
	return strings.Map(ck.enc, pt)
}

func (ck ckey) decipher(ct string) string {
	return strings.Map(ck.dec, ct)
}

func main() {

	pt := "The five boxing wizards jump quickly"

	for _, key := range []int{0,1,7,25,26} {
		ck, ok := newCaeser(key)
		if !ok {
			fmt.Println("key", key, "invalid")
			continue
		}
		ct := ck.encipher(pt)
		fmt.Println("key ", key)
		fmt.Println(" Enciphered: ", ct)
		fmt.Println("  Deciphered: ", ck.decipher(ct))
	}

}
