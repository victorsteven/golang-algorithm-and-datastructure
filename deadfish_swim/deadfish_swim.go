package deadfish_swim

// Write a simple parser that will parse and run Deadfish.
//
//Deadfish has 4 commands, each 1 character long:
//
//i increments the value (initially 0)
//d decrements the value
//s squares the value
//o outputs the value into the return array
//Invalid characters should be ignored.
//
//Parse("iiisdoso") == []int{8, 64}

func Parse(data string) []int {
	output := []int{}

	var val int
	for _, s := range data {
		switch string(s) {
		case "i":
			val++
		case "d":
			val--
		case "s":
			val = val * val
		case "o":
			output = append(output, val)
		}
	}
	return output
}
