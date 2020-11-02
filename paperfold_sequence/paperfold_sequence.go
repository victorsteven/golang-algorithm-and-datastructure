package paperfold_sequence

//Wikipedia: The regular paperfolding sequence, also known as the dragon curve sequence, is an infinite automatic sequence of 0s and 1s defined as the limit of the following process:
//
//1
//1 1 0
//1 1 0 1 1 0 0
//1 1 0 1 1 0 0 1 1 1 0 0 1 0 0
//
//At each stage an alternating sequence of 1s and 0s is inserted between the terms of the previous sequence.
//
//Define a generator PaperFold that sequentially generates the values of this sequence.
//
//It will be tested for up to 1 000 000 values.

func PaperFold(ch chan <- int) {
	v := 1
	i := 1
	for true {
		v = i
		for v % 2 == 0 {
			v /= 2
		}
		if v % 4 == 1 {
			ch <- 1
		} else {
			ch <- 0
		}
		i++
	}
}

