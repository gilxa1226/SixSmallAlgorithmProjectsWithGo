package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const num_rows = 8
const num_cols = num_rows

// Whether we want an open or closed tour.
const require_closed_tour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var move_offsets []Offset

var num_calls int64

func initialize_offsets() {
	move_offsets = []Offset{
		Offset{-2, -1},
		Offset{-1, -2},
		Offset{+2, -1},
		Offset{+1, -2},
		Offset{-2, +1},
		Offset{-1, +2},
		Offset{+2, +1},
		Offset{+1, +2},
	}
}

func make_board(num_rows int, num_cols int) [][]int {
	tmp := make([][]int, num_rows)
	for r := 0; r < num_rows; r++ {
		tmp[r] = make([]int, num_cols)
		for c := 0; c < num_cols; c++ {
			tmp[r][c] = -1
		}
	}
	return tmp
}

func dump_board(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%02d ", val)
		}
		fmt.Println()
	}
}

func onboard(c int, r int, num_cols int, num_rows int) bool {
	return (c >= 0 && c < num_cols) && (r >= 0 && r < num_rows)
}

func find_tour(board [][]int, num_rows, num_cols, cur_row, cur_col, num_visited int) bool {
	num_calls += 1

	if num_visited == num_rows*num_cols {
		if require_closed_tour == false {
			return true
		} else {
			for _, row := range board {
				for _, val := range row {
					if val == 0 {
						return true
					}
				}
			}
			return false
		}
	} else {
		for _, move := range move_offsets {
			newc := cur_col + move.dc
			newr := cur_row + move.dr
			if !onboard(newc, newr, num_cols, num_rows) {
				continue
			}

			if board[newr][newc] != -1 {
				continue
			}

			board[newr][newc] = num_visited
			if find_tour(board, num_rows, num_cols, newr, newc, num_visited+1) {
				return true
			}
			board[newr][newc] = -1
		}
		return false
	}
}

func main() {
	num_calls = 0

	// Initialize the move offsets.
	initialize_offsets()

	// Create the blank board.
	board := make_board(num_rows, num_cols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if find_tour(board, num_rows, num_cols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dump_board(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", num_calls)
}
