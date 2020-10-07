package baseE91

//BasE91 is a method for encoding binary as ASCII characters. It is more efficient than Base64 and needs 91 characters to represent the encoded data.
//
//The following ASCII charakters are used:
//
//'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
//'!#$%&()*+,./:;<=>?@[]^_`{|}~"'
//Create two functions that encode strings to basE91 string and decodes the other way round.
//
//b91encode('test') = 'fPNKd'
//b91decode('fPNKd') = 'test'
//
//b91decode('>OwJh>Io0Tv!8PE') = 'Hello World!'
//b91encode('Hello World!') = '>OwJh>Io0Tv!8PE'


var enctab = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'")

// Decoding table maps all the characters back to their integer values
var dectab = map[byte]byte{
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
	'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
	'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
	'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
	'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
	'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
	'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
	'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '!': 62, '#': 63,
	'$': 64, '%': 65, '&': 66, '(': 67, ')': 68, '*': 69, '+': 70, ',': 71,
	'.': 72, '/': 73, ':': 74, ';': 75, '<': 76, '=': 77, '>': 78, '?': 79,
	'@': 80, '[': 81, ']': 82, '^': 83, '_': 84, '`': 85, '{': 86, '|': 87,
	'}': 88, '~': 89, '\'': 90,
}

func Encode(d []byte) []byte {

	var n, b uint
	var o []byte

	for i := 0; i < len(d); i++ {
		b |= uint(d[i]) << n
		n += 8
		if n > 13 {
			v := b & 8191
			if v > 88 {
				b >>= 13
				n -= 13
			} else {
				v = b & 16383
				b >>= 14
				n -= 14
			}
			o = append(o, enctab[v%91], enctab[v/91])
		}
	}

	if n > 0 {

		o = append(o, enctab[b%91])
		if n > 7 || b > 90 {
			o = append(o, enctab[b/91])
		}
	}
	return o
}

func Decode(d []byte) []byte {
	var b, n uint
	var o []byte
	v := -1

	for i := 0; i < len(d); i++ {
		c, ok := dectab[d[i]]
		if !ok {
			continue
		}
		if v < 0 {
			v = int(c)
		} else {
			v += int(c) * 91
			b |= uint(v) << n
			if v&8191 > 88 {
				n += 13
			} else {
				n += 14
			}
			o = append(o, byte(b&255))
			b >>= 8
			n -= 8
			for n > 7 {
				o = append(o, byte(b&255))
				b >>= 8
				n -= 8
			}
			v = -1
		}
	}
	if v+1 > 0 {
		o = append(o, byte((b|uint(v)<<n)&255))
	}
	return o
}


func Encode2(ib []byte) []byte {
	ebq, en := 0, uint(0)
	r := make([]byte, 0, 2*len(ib))
	for i := range ib {
		ebq |= int(ib[i]) << en
		en += 8
		if en > 13 {
			ev := ebq & 8191
			if ev > 88 {
				ebq >>= 13
				en -= 13
			} else {
				ev = ebq & 16383
				ebq >>= 14
				en -= 14
			}
			r = append(r, enctab2[ev%91])
			r = append(r, enctab2[ev/91])
		}
	}
	if en > 0 {
		r = append(r, enctab2[ebq%91])
		if en > 7 || ebq > 90 {
			r = append(r, enctab2[ebq/91])
		}
	}
	return r
}
func Decode2(ib []byte) []byte {
	dbq, dn := 0, uint(0)
	var dv int = -1
	r := []byte{}
	for i := range ib {
		if dectab2[ib[i]] == 255 {
			continue
		}
		if dv == -1 {
			dv = int(dectab2[ib[i]])
		} else {
			dv += int(dectab2[ib[i]]) * 91
			dbq |= dv << dn
			dn += 13
			if dv&8191 <= 88 {
				dn++
			}
			for {
				r = append(r, byte(dbq))
				dbq >>= 8
				dn -= 8
				if dn <= 7 {
					break
				}
			}
			dv = -1
		}
	}
	if dv != -1 {
		r = append(r, (byte)(dbq|dv<<dn))
	}
	return r
}

func init() {
	for i := range dectab2 {
		dectab2[i] = 255
	}
	for i := 0; i < 91; i++ {
		dectab2[enctab2[i]] = byte(i)
	}
}

var enctab2 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'")
var dectab2 = [256]byte{}