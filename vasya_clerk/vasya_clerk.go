package vasya_clerk

//The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line. Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.
//
//Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.
//
//Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?
//
//Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.
//
//Examples:
//Tickets([]int{25, 25, 50}) // => YES
//Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to give change to 100 dollars
//Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two bills of 25 from one of 50)

func Tickets(peopleInLine []int) string {

	t, f := 0, 0

	for _, bill := range peopleInLine {
		switch bill {
		case 25:
			t++
		case 50:
			f++
			t--
		case 100:
			if f > 0 {
				f--
				t--
			} else {
				t -= 3
			}
		}
		if t < 0 || f < 0 {
			return "NO"
		}
	}

	return "YES"
}

func Tickets2(peopleInLine []int) string {
	bank := map[int]int{25: 0, 50: 0, 100: 0}

	for _, bill := range peopleInLine {
		bank[bill]++
		switch bill {
		case 50:
			if bank[25] == 0 {
				return "NO"
			}
			bank[25]--
		case 100:
			if bank[50] > 0 && bank[25] > 0 {
				bank[25]--
				bank[50]--
			} else if bank[25] >= 3 {
				bank[25] -= 3
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}

func Tickets3(peopleInLine []int) string {

	if peopleInLine[0] > 25 {
		return "NO"
	}
	wallet := make(map[int]int)

	for _, v := range peopleInLine {
		change := v - 25
		if change == 25 {
			if wallet[change] <= 0 {
				return "NO"
			}
			wallet[25]--
		} else if change == 75 {
			if wallet[25] >= 1 && wallet[50] >= 1 {
				wallet[25]--
				wallet[50]--
			} else if wallet[25] >= 3 {
				wallet[25] -= 3
			} else {
				return "NO"
			}
		}
		wallet[v]++
	}
	return "YES"
}
