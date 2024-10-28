package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"slices"
	"strconv"
)

const sudoku_csv = "sudoku.csv"

// panics if input file is invalid
func getProblem(filepath string) iter.Seq2[[]int, []int] {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	return func(yield func([]int, []int) bool) {
		for {
			problem, err := reader.Read()
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				panic("some read error")
			}
			if problem[0] == "quizzes" {
				// skip header
				continue
			}
			quiz, err := stringToIntArray(problem[0])
			if err != nil {
				panic(err)
			}
			answer, err := stringToIntArray(problem[1])
			if err != nil {
				panic(err)
			}
			if !yield(quiz, answer) {
				return
			}
		}
	}
}

func stringToIntArray(s string) ([]int, error) {
	ints := make([]int, len(s))
	for i, char := range s {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}
	return ints, nil
}

func solve(sudoku []int) []int {
	if sudoku == nil {
		return nil
	}
	for col, row, ok := findCandidatePos(sudoku); ok; {
		for num := 1; num <= 9; num++ {
			if !check(sudoku, num, col, row) {
				continue
			}
			newSudoku := make([]int, len(sudoku))
			copy(newSudoku, sudoku)
			newSudoku[row*9+col] = num
			if result := solve(newSudoku); result != nil {
				return result
			}
		}
		return nil
	}
	return sudoku
}

func findCandidatePos(sudoku []int) (col, row int, ok bool) {
	for i, num := range sudoku {
		if num == 0 {
			return i % 9, i / 9, true
		}
	}
	return -1, -1, false
}

func check(sudoku []int, num int, col, row int) bool {
	return checkRow(sudoku, num, row) && checkCol(sudoku, num, col) && checkSquare(sudoku, num, col, row)
}

func checkRow(sudoku []int, num int, row int) bool {
	for col := range 9 {
		if sudoku[row*9+col] == num {
			return false
		}
	}
	return true
}

func checkCol(sudoku []int, num int, col int) bool {
	for row := range 9 {
		if sudoku[row*9+col] == num {
			return false
		}
	}
	return true
}

func checkSquare(sudoku []int, num int, col, row int) bool {
	pos := 9*3*(row/3) + 3*(col/3)
	for range 3 {
		for squareNum := range slices.Values(sudoku[pos : pos+3]) {
			if squareNum == num {
				return false
			}
		}
		pos += 9
	}
	return true
}

func main() {
	for sudoku, answer := range getProblem(sudoku_csv) {
		guess := solve(sudoku)

		if !slices.Equal(guess, answer) {
			panic("wrong answer")
		}
	}
}

func prettyPrint(sudoku []int) {
	for i, num := range sudoku {
		fmt.Printf("%v ", num)
		if i%9 == 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}
