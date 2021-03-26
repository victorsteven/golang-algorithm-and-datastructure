package main

import (
	"math"
)

//My friend John likes to go to the cinema. He can choose between system A and system B.
//
//System A : he buys a ticket (15 dollars) every time
//System B : he buys a card (500 dollars) and a first ticket for 0.90 times the ticket price,
//then for each additional ticket he pays 0.90 times the price paid for the previous ticket.
//#Example: If John goes to the cinema 3 times:
//
//System A : 15 * 3 = 45
//System B : 500 + 15 * 0.90 + (15 * 0.90) * 0.90 + (15 * 0.90 * 0.90) * 0.90 ( = 536.5849999999999, no rounding for each ticket)
//John wants to know how many times he must go to the cinema so that the final result of System B, when rounded up to the next dollar, will be cheaper than System A.
//
//The function movie has 3 parameters: card (price of the card), ticket (normal price of a ticket), perc (fraction of what he paid for the previous ticket) and returns the first n such that
//
//ceil(price of System B) < price of System A.
//More examples:
//
//movie(500, 15, 0.9) should return 43
//    (with card the total price is 634, with tickets 645)
//movie(100, 10, 0.95) should return 24
//    (with card the total price is 235, with tickets 240)

func Movie(card, ticket int, perc float64) (count int) {

	var systemA float64

	systemB := float64(card)

	for systemA <= math.Ceil(systemB) {
		count++
		systemA = float64(ticket * count)
		systemB += float64(ticket) * math.Pow(perc, float64(count))
	}

	return count
}
