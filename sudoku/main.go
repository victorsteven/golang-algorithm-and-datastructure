package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	board := parseInput(os.Args[1])

	printBoard(board)

	if backtrack(&board) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(board)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}


//Sudoku Background
//Sudoku is a game played on a 9x9 grid. The goal of the game is to fill all cells of the grid with digits from 1 to 9, so that each column, each row, and each of the nine 3x3 sub-grids (also known as blocks) contain all of the digits from 1 to 9.
//(More info at: http://en.wikipedia.org/wiki/Sudoku)
//
//Sudoku Solution Validator
//Write a function validSolution/ValidateSolution/valid_solution() that accepts a 2D array representing a Sudoku board, and returns true if it is a valid solution, or false otherwise. The cells of the sudoku board may also contain 0's, which will represent empty cells. Boards containing one or more zeroes are considered to be invalid solutions.
//
//The board is always 9 cells by 9 cells, and every cell only contains integers from 0 to 9.


//validSolution([
//[5, 3, 4, 6, 7, 8, 9, 1, 2],
//[6, 7, 2, 1, 9, 5, 3, 4, 8],
//[1, 9, 8, 3, 4, 2, 5, 6, 7],
//[8, 5, 9, 7, 6, 1, 4, 2, 3],
//[4, 2, 6, 8, 5, 3, 7, 9, 1],
//[7, 1, 3, 9, 2, 4, 8, 5, 6],
//[9, 6, 1, 5, 3, 7, 2, 8, 4],
//[2, 8, 7, 4, 1, 9, 6, 3, 5],
//[3, 4, 5, 2, 8, 6, 1, 7, 9]
//]); // => true
//validSolution([
//[5, 3, 4, 6, 7, 8, 9, 1, 2],
//[6, 7, 2, 1, 9, 0, 3, 4, 8],
//[1, 0, 0, 3, 4, 2, 5, 6, 0],
//[8, 5, 9, 7, 6, 1, 0, 2, 0],
//[4, 2, 6, 8, 5, 3, 7, 9, 1],
//[7, 1, 3, 9, 2, 4, 8, 5, 6],
//[9, 0, 1, 5, 3, 7, 2, 1, 4],
//[2, 8, 7, 4, 1, 9, 6, 3, 5],
//[3, 0, 0, 4, 8, 1, 1, 7, 9]
//]); // => false

//func main() {
	//fmt.Println([]int{1:3})

	//numbers := [5]int{1, 2, 3, 4, 5}
	//fmt.Println((&numbers)[1:4]) // [2, 3]
	//
	//A := [][]int{
	//	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	//	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	//	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	//	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	//	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	//	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	//	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	//	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	//	{3, 4, 5, 2, 8, 6, 1, 7, 9},
	//}
	//for i := 0; i < len(A); i++ {
	//	fmt.Println((A[i])[1:4])
	//	//ans := []int{}
	//	//ans := append(ans, (A[i])[1:4]) // [2, 3]
	//}
	//
	//B := [][]int{
	//	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	//	{6, 7, 2, 1, 9, 0, 3, 4, 8},
	//	{1, 0, 0, 3, 4, 2, 5, 6, 0},
	//	{8, 5, 9, 7, 6, 1, 0, 2, 0},
	//	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	//	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	//	{9, 0, 1, 5, 3, 7, 2, 1, 4},
	//	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	//	{3, 0, 0, 4, 8, 1, 1, 7, 9},
	//}

	//fmt.Println(ValidateSolution(B))
	//fmt.Println(ValidateSolution(A))

	//fmt.Println(getSelection(A))

//	board := parseInput(os.Args[1])
//
//	printBoard(board)
//
//	if backtrack(&board) {
//		fmt.Println("The Sudoku was solved successfully:")
//		printBoard(board)
//	} else {
//		fmt.Printf("The Sudoku can't be solved.")
//	}
//}



//func ValidateSolution(m [][]int) bool {
//	// Your code here
//	//ans := [][]int{}
//	outerLength := len(m)
//	innerLength := len(m[0]) //choose any the of the rows and get the length
//	for i := 0; i < outerLength; i++ {
//		for j := 0; j < innerLength; j++ {
//			if m[i][j] <= 0  {
//				return true
//			}
//		}
//	}
//	return false
//}

//func getSelection(m, n int, board [][]int) {
//	//arr := make([]int, 0)
//	for i := 3*m; i < 3 * (m+1); i++ {
//		//newArray := []int{}
//		//fmt.Println([]int{3*n:3})
//		//newArray = append(newArray, (board[i])[1:4])
//
//		board = append(board, (board[i])[1:4])
//		//arr = append(arr, )
//	}
//}

