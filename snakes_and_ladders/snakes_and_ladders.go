package snakes_and_ladders

import "fmt"

//Snakes and Ladders is an ancient Indian board game regarded today as a worldwide classic. It is played between two or more players on a gameboard having numbered, gridded squares. A number of "ladders" and "snakes" are pictured on the board, each connecting two specific board squares. (Source Wikipedia)

//Your task is to make a simple class called SnakesLadders. The test cases will call the method play(die1, die2) independantly of the state of the game or the player turn. The variables die1 and die2 are the die thrown in a turn and are both integers between 1 and 6. The player will move the sum of die1 and die2.

//Rules
//1.  There are two players and both start off the board on square 0.
//
//2.  Player 1 starts and alternates with player 2.
//
//3.  You follow the numbers up the board in order 1=>100
//
//4.  If the value of both die are the same then that player will have another go.
//
//5.  Climb up ladders. The ladders on the game board allow you to move upwards and get ahead faster. If you land exactly on a square that shows an image of the bottom of a ladder, then you may move the player all the way up to the square at the top of the ladder. (even if you roll a double).
//
//6.  Slide down snakes. Snakes move you back on the board because you have to slide down them. If you land exactly at the top of a snake, slide move the player all the way to the square at the bottom of the snake or chute. (even if you roll a double).
//
//7.  Land exactly on the last square to win. The first person to reach the highest square on the board wins. But there's a twist! If you roll too high, your player "bounces" off the last square and moves back. You can only win by rolling the exact number needed to land on the last square. For example, if you are on square 98 and roll a five, move your game piece to 100 (two moves), then "bounce" back to 99, 98, 97 (three, four then five moves.)
//
//8.  If the Player rolled a double and lands on the finish square “100” without any remaining moves then the Player wins the game and does not have to roll again.

var m = map[int]int{2: 38, 7: 14, 8: 31, 15: 26, 16: 6, 21: 42, 28: 84, 36: 44, 46: 25, 49: 11, 51: 67, 62: 19, 64: 60, 71: 91, 74: 53, 78: 98, 87: 94, 89: 68, 92: 88, 95: 75, 99: 80}

// SnakesLadders - Stucture that game is played on
type SnakesLadders struct {
	// Define Game Struct
	player [2]int
	turn int
}

// NewGame - Method that starts a new game for the SnakesLadders struct
func (g *SnakesLadders) NewGame() {
	// your code here
	g.turn, g.player[0], g.player[1] = 0, 0, 0
}

// Play - Method that makes turn givem a doce roll from die1 and die2 for the SnakesLadders struct
// Player that dice is for should automatically be determined based on rules
func (g *SnakesLadders) Play(die1, die2 int) string {
	// your code here
	if g.turn == -1 {
		return "Game over!"
	}

	pos := g.player[g.turn] + die1 + die2

	if pos > 100 {
		pos = 100 - pos % 100
	}

	if m[pos] != 0 {
		pos = m[pos]
	}

	g.player[g.turn] = pos

	s := fmt.Sprintf("Player %d is on square %d", g.turn + 1, pos)

	if pos == 100 {
		s = fmt.Sprintf("Player %d wins!", g.turn + 1)
		g.turn = -1
	} else {
		if die1 != die2 {
			g.turn ^= 1
		}
	}
	return s
}

// It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 1,"Player 1 is on square 38"),func() {Expect(game.Play(1, 1)).To(Equal("Player 1 is on square 38"))} )
//    It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 5,"Player 1 is on square 44"),func() {Expect(game.Play(1, 5)).To(Equal("Player 1 is on square 44"))} )
//    It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",6, 2,"Player 2 is on square 31"),func() {Expect(game.Play(6, 2)).To(Equal("Player 2 is on square 31"))} )
//    It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 1,"Player 1 is on square 25"),func() {Expect(game.Play(1, 1)).To(Equal("Player 1 is on square 25"))} )


