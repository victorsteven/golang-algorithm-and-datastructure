package __by_4_skyscrappers

import (
	"fmt"
	"strconv"
)

//In a grid of 4 by 4 squares you want to place a skyscraper in each square with only some clues:
//
//The height of the skyscrapers is between 1 and 4
//No two skyscrapers in a row or column may have the same number of floors
//A clue is the number of skyscrapers that you can see in a row or column from the outside
//Higher skyscrapers block the view of lower skyscrapers located behind them
//
//Can you write a program that can solve this puzzle?
//
//Example:
//
//To understand how the puzzle works, this is an example of a row with 2 clues. Seen from the left side there are 4 buildings visible while seen from the right side only 1:


//If no clue is available, add value `0`
//Each puzzle has only one possible solution
//`SolvePuzzle()` returns matrix `int[][]`. The first indexer is for the row, the second indexer for the column. (Python: returns 4-tuple of 4-tuples, Ruby: 4-Array of 4-Arrays)

func SolvePuzzle(clues []int) [][]int {
	ans := make([][]int, 4)
	ans[0] = []int{1,2,3,4}
	for t1 := true; t1; t1 = next(ans[0]) {
		ans[1] = []int{1,2,3,4}
		for t2 := true; t2; t2 = next(ans[1]) {
			ans[2] = []int{1,2,3,4}
			for t3 := true; t3; t3 = next(ans[2]) {
				ans[3] = []int{1,2,3,4}
				for t4 := true; t4; t4 = next(ans[3]) {
					if check(ans, clues) {
						return ans
					}
				}
			}
		}
	}

	// never rich
	return nil
}

func check(puzzle [][]int, clues []int) bool {
	// check 0,1,2,3
	for i := 0; i < 4; i++ {
		cnt := seen([]int{puzzle[0][i], puzzle[1][i], puzzle[2][i], puzzle[3][i]})
		if cnt > 5 {return false}
		if clues[i] == 0 {continue}
		if cnt != clues[i] {return false}
	}
	// check 4-7
	for i := 0; i < 4; i++ {
		cnt := seen([]int{puzzle[i][3], puzzle[i][2], puzzle[i][1], puzzle[i][0]})
		if cnt > 5 {return false}
		if clues[4+i] == 0 {continue}
		if cnt != clues[4+i] {return false}
	}

	// check 8-11
	for i := 0; i < 4; i++ {
		cnt := seen([]int{puzzle[3][3-i], puzzle[2][3-i], puzzle[1][3-i], puzzle[0][3-i]})
		if cnt > 5 {return false}
		if clues[8+i] == 0 {continue}
		if cnt != clues[8+i] {return false}
	}

	// check 12-15
	for i := 0; i < 4; i++ {
		cnt := seen([]int{puzzle[3-i][0], puzzle[3-i][1], puzzle[3-i][2], puzzle[3-i][3]})
		if cnt > 5 {return false}
		if clues[12+i] == 0 {continue}
		if cnt != clues[12+i] {return false}
	}

	return true
}

// returns number of skyscabes seen in a row
func seen(row []int) int {
	m := row[0]
	used := make([]bool, 5)
	used[m] = true
	c := 1
	for i := 1; i < 4; i++ {
		if used[row[i]] { return 7 }
		used[row[i]] = true
		if row[i] > m {
			m = row[i]
			c++
		}
	}
	return c
}

func next(p []int) bool {
	n := len(p)
	i := n - 2
	for i >= 0 {
		if p[i] < p[i + 1] { break}
		i--
	}
	if i < 0 {return false}
	j := n - 1
	for j > i {
		if p[j] > p[i] {break}
		j--
	}
	if j == i { return false}
	p[i], p[j] = p[j], p[i]
	reverse(p[i + 1:])

	return true
}

func reverse(p []int) {
	n := len(p)
	for i := 0; i < n/2; i++ {
		p[i], p[n - 1 - i] = p[n - 1 - i], p[i]
	}
}



////////////////////////////////////////////////////////////////
type placement struct {
	row int
	col int
}

var ordrNums [24]int
var clueToPlacement []placement

func SolvePuzzle2(clues []int) [][]int {
	clueToPlacement = []placement{
		placement{0, 0}, placement{0, 1}, placement{0, 2}, placement{0, 3},
		placement{0, 3}, placement{1, 3}, placement{2, 3}, placement{3, 3},
		placement{3, 3}, placement{3, 2}, placement{3, 1}, placement{3, 0},
		placement{3, 0}, placement{2, 0}, placement{1, 0}, placement{0, 0}}

	ordrNums = [24]int{1234, 1243, 1324, 1342, 1423, 1432, 2134, 2143, 2314, 2341, 2413, 2431, 3124, 3142, 3214, 3241, 3412, 3421, 4123, 4132, 4213, 4231, 4312, 4321}

	knownBoard := [4][4]int{}
	// do 4s and 1s
	for clue := 0; clue < 16; clue++ {
		row, col := GetAt(clue)
		if clues[clue] == 1 {
			knownBoard[row][col] = 4
		} else if clues[clue] == 4 {
			knownBoard[row][col] = 1

			switch (clue) / 4 {
			case 0:
				knownBoard[row+1][col], knownBoard[row+2][col], knownBoard[row+3][col] = 2, 3, 4
				break
			case 1:
				knownBoard[row][col-1], knownBoard[row][col-2], knownBoard[row][col-3] = 2, 3, 4
				break
			case 2:
				knownBoard[row-1][col], knownBoard[row-2][col], knownBoard[row-3][col] = 2, 3, 4
				break
			case 3:
				knownBoard[row][col+1], knownBoard[row][col+2], knownBoard[row][col+3] = 2, 3, 4
				break
			}
		}
	}

	board := knownBoard
	//find all the options that can fit in each col
	posibilitiesForCols := [4][]int{}
	for i := 0; i < 4; i++ {
		for num := range ordrNums {
			fits := true
			digits := strconv.Itoa(ordrNums[num])
			for row := 0; row < 4; row++ {
				dig, err := strconv.ParseInt(string(digits[row]), 10, 0)
				if err != nil {
					fmt.Print("ohoh!")
				}
				if board[row][i] != 0 && board[row][i] != int(dig) {
					fits = false
					break
				}
			}
			if fits {
				posibilitiesForCols[i] = append(posibilitiesForCols[i], ordrNums[num])
			}
		}
	}

	fboard := [4][4]int{}
	for _, col0 := range posibilitiesForCols[0] {
		fboard = InsertCol(fboard, col0, 0)
		for _, col1 := range posibilitiesForCols[1] {
			fboard = InsertCol(fboard, col1, 1)
			for _, col2 := range posibilitiesForCols[2] {
				fboard = InsertCol(fboard, col2, 2)
				for _, col3 := range posibilitiesForCols[3] {
					fboard = InsertCol(fboard, col3, 3)
					if CheckIfAnswer(fboard, clues) {
						return TurnToSlice(fboard)
					}
				}
			}
		}
	}
	return nil
}
func InsertCol(board [4][4]int, col int, colInt int) [4][4]int {
	for i := 0; i < 4; i++ {
		digits := strconv.Itoa(col)
		dig, err := strconv.ParseInt(string(digits[i]), 10, 0)
		if err != nil {
			fmt.Print("ohoh!")
		}
		board[i][colInt] = int(dig)
	}
	return board
}
func GetAt(clue int) (int, int) {
	return clueToPlacement[clue].row, clueToPlacement[clue].col
}
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func CheckIfAnswer(board [4][4]int, clues []int) bool {
	//?check if rows valid
	for row := 0; row < 4; row++ {
		a := []int{}
		for col := 0; col < 4; col++ {
			if IntInSlice(board[row][col], a) {
				return false
			}
			a = append(a, board[row][col])
		}
	}
	//?check if cols valid
	for col := 0; col < 4; col++ {
		a := []int{}
		for row := 0; row < 4; row++ {
			if IntInSlice(board[row][col], a) {
				return false
			}
			a = append(a, board[row][col])
		}
	}

	//?check if right
	for c := 0; c < len(clues); c++ {
		count := 1
		switch (c) / 4 {
		case 0:
			count = VisableAmount(c, board, 1, 0)
			break
		case 1:
			count = VisableAmount(c, board, 0, -1)
			break
		case 2:
			count = VisableAmount(c, board, -1, 0)
			break
		case 3:
			count = VisableAmount(c, board, 0, +1)
			break
		}
		if count != clues[c] && clues[c] != 0 {
			return false
		}
	}
	return true
}
func VisableAmount(clue int, board [4][4]int, addToR int, addToC int) int {
	row, col := GetAt(clue)
	hight := board[row][col]
	count := 1
	if board[row+addToR][col+addToC] > hight {
		hight = board[row+addToR][col+addToC]
		count++
	}
	if board[row+(addToR*2)][col+(addToC*2)] > hight {
		hight = board[row+(addToR*2)][col+(addToC*2)]
		count++
	}
	if board[row+(addToR*3)][col+(addToC*3)] > hight {
		hight = board[row+(addToR*3)][col+(addToC*3)]
		count++
	}
	return count
}
func TurnToSlice(arr [4][4]int) [][]int {
	var result [4][]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i] = append(result[i], arr[i][j])
		}
	}
	return result[:]
}




///////////////////////////////////////////////////////////////////
func perm(xn []int, i int, out *[][]int) {
	if i == len(xn)-1 {
		temp := make([]int, len(xn))
		copy(temp, xn)
		*out = append(*out, temp)
		return
	}
	for j := i; j < len(xn); j++ {
		xn[i], xn[j] = xn[j], xn[i]
		perm(xn, i+1, out)
		xn[i], xn[j] = xn[j], xn[i]
	}

}

func verticalUnique(xn [][]int) bool {
	for i := 0; i < len(xn); i++ {
		col := 0
		for j := 0; j < len(xn[0]); j++ {
			col ^= 1 << uint(xn[j][i])
		}
		if col != 30 {
			return false
		}
	}
	return true
}

func checkDown(xn [][]int, clues []int) bool {
	for i := 0; i < 4; i++ {
		if clues[i] == 0 {
			continue
		}
		max, seen := xn[0][i], 1
		for j := 1; j < 4; j++ {
			if xn[j][i] > max {
				seen++
				max = xn[j][i]
			}
		}
		if seen != clues[i] {
			return false
		}
	}
	return true
}

func checkRight(xn [][]int, clues []int) bool {
	for i := 4; i < 8; i++ {
		if clues[i] == 0 {
			continue
		}
		max, seen := xn[i%4][3], 1
		for j := 2; j >= 0; j-- {
			if xn[i%4][j] > max {
				seen++
				max = xn[i%4][j]
			}
		}
		if seen != clues[i] {
			return false
		}
	}
	return true
}

func checkUp(xn [][]int, clues []int) bool {
	for i := 8; i < 12; i++ {
		if clues[i] == 0 {
			continue
		}
		max, seen := xn[3][3-i%4], 1
		for j := 2; j >= 0; j-- {
			if xn[j][3-i%4] > max {
				seen++
				max = xn[j][3-i%4]
			}
		}
		if seen != clues[i] {
			return false
		}
	}
	return true
}

func checkLeft(xn [][]int, clues []int) bool {
	for i := 12; i < 16; i++ {
		if clues[i] == 0 {
			continue
		}
		max, seen := xn[3-i%4][0], 1
		for j := 1; j < 4; j++ {
			if xn[3-i%4][j] > max {
				seen++
				max = xn[3-i%4][j]
			}
		}
		if seen != clues[i] {
			return false
		}
	}
	return true
}

func SolvePuzzle3(clues []int) [][]int {
	n := 24
	grid := make([][]int, 4)
	p := make([][]int, 0, n)
	xn := []int{
		1, 2, 3, 4,
	}
	perm(xn, 0, &p)

	for i := 0; i < n; i++ {
		grid[0] = p[i]
		for j := 0; j < n; j++ {
			grid[1] = p[j]
			for k := 0; k < n; k++ {
				grid[2] = p[k]
				for l := 0; l < n; l++ {
					grid[3] = p[l]
					if verticalUnique(grid) &&
						checkDown(grid, clues) &&
						checkRight(grid, clues) &&
						checkUp(grid, clues) &&
						checkLeft(grid, clues) {
						goto end
					}
				}
			}
		}
	}
end:
	return grid
}