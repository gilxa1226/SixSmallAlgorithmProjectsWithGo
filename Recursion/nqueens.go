package main

import (
	"fmt"
	"time"
)

const QUEEN = "Q"

var attack_counts [][]int

func make_board(num_rows int) [][]string {
	tmp := make([][]string, num_rows)
	for idx, _ := range tmp {
		tmp[idx] = make([]string, num_rows)
		for ydx, _ := range tmp[idx] {
			tmp[idx][ydx] = "."
		}
	}
	return tmp
}

func dump_board(board [][]string) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%s ", board[r][c])
		}
		fmt.Println()
	}
}

func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
	num_queens := 0
	row, col := r0, c0
	for {
		if board[row][col] == QUEEN {
			num_queens++
		}
		row += dr
		col += dc
		if row >= num_rows || col >= num_rows {
			break
		}
	}

	return (num_queens < 2)
}

func board_is_legal(board [][]string, num_rows int) bool {

	// rows + diagonals
	for r := 0; r < num_rows; r++ {
		if !series_is_legal(board, num_rows, r, 0, 0, 1) {
			return false
		}
		if !series_is_legal(board, num_rows, r, 0, 1, 1) {
			return false
		}
		if !series_is_legal(board, num_rows, r, num_rows-1, 1, -1) {
			return false
		}
	}

	// columns + diagonals
	for c := 0; c < num_rows; c++ {
		if !series_is_legal(board, num_rows, 0, c, 1, 0) {
			return false
		}

		if !series_is_legal(board, num_rows, 0, c, 1, 1) {
			return false
		}

		if !series_is_legal(board, num_rows, num_rows-1, c, -1, 1) {
			return false
		}
	}

	return true
}

func board_is_a_solution(board [][]string, num_rows int) bool {

	num_queens := 0

	if !board_is_legal(board, num_rows) {
		return false
	}

	for r := 0; r < num_rows; r++ {
		for c := 0; c < num_rows; c++ {
			if board[r][c] == QUEEN {
				num_queens++
			}
		}
	}

	if num_queens == num_rows {
		return true
	} else {
		return false
	}

}

func place_queens_1(board [][]string, num_rows, r, c int) bool {

	if r >= num_rows {
		return board_is_a_solution(board, num_rows)
	}

	next_c := c + 1
	next_r := r

	if next_c >= num_rows {
		next_r += 1
		next_c = 0
	}

	test := place_queens_1(board, num_rows, next_r, next_c)

	if test {
		return true
	}

	board[r][c] = QUEEN

	test = place_queens_1(board, num_rows, next_r, next_c)

	if test {
		return true
	}

	board[r][c] = "."

	return false
}

func place_queens_2(board [][]string, num_rows, num_placed, r, c int) bool {

	if r >= num_rows {
		return board_is_a_solution(board, num_rows)
	}

	if num_placed == num_rows {
		return board_is_a_solution(board, num_rows)
	}

	next_c := c + 1
	next_r := r

	if next_c >= num_rows {
		next_r += 1
		next_c = 0
	}

	test := place_queens_2(board, num_rows, num_placed, next_r, next_c)

	if test {
		return true
	}

	board[r][c] = QUEEN

	test = place_queens_2(board, num_rows, num_placed+1, next_r, next_c)

	if test {
		return true
	}

	board[r][c] = "."

	return false
}

func init_attack_counts(num_rows int) {
	attack_counts = make([][]int, num_rows)
	for r, _ := range attack_counts {
		attack_counts[r] = make([]int, num_rows)
		for c, _ := range attack_counts[r] {
			attack_counts[r][c] = 0
		}
	}
}

func adjust_attack_counts(r int, c int, inc int, num_rows int) {
	// adjust col
	for col := 0; col < num_rows; col++ {
		attack_counts[r][col] += inc
	}

	// adjust row
	for row := 0; row < num_rows; row++ {
		attack_counts[row][c] += inc
	}

	// adjust both diagonals
	// up + left
	row, col := r, c
	for {
		row--
		col--
		if row >= 0 && col >= 0 {
			attack_counts[row][col] += inc
		} else {
			break
		}
	}

	// up + right
	row, col = r, c
	for {
		row--
		col++
		if row >= 0 && col < num_rows {
			attack_counts[row][col] += inc
		} else {
			break
		}
	}

	// down + left
	row, col = r, c
	for {
		row++
		col--
		if row < num_rows && col > 0 {
			attack_counts[row][col] += inc
		} else {
			break
		}
	}

	// down + right
	row, col = r, c
	for {
		row++
		col++
		if row < num_rows && col < num_rows {
			attack_counts[row][col] += inc
		} else {
			break
		}
	}
}
func place_queens_3(board [][]string, num_rows, num_placed, r, c int) bool {

	if r >= num_rows {
		return board_is_a_solution(board, num_rows)
	}

	if num_placed == num_rows {
		return board_is_a_solution(board, num_rows)
	}

	next_c := c + 1
	next_r := r

	if next_c >= num_rows {
		next_r += 1
		next_c = 0
	}

	test := place_queens_3(board, num_rows, num_placed, next_r, next_c)

	if test {
		return true
	}

	if attack_counts[r][c] == 0 {
		board[r][c] = QUEEN
		adjust_attack_counts(r, c, +1, num_rows)
		test = place_queens_3(board, num_rows, num_placed+1, next_r, next_c)

		if test {
			return true
		}

		board[r][c] = "."
		adjust_attack_counts(r, c, -1, num_rows)
	}

	return false
}

func main() {
	const num_rows = 13
	board := make_board(num_rows)

	start := time.Now()
	//success := place_queens_1(board, num_rows, 0, 0)
	//success := place_queens_2(board, num_rows, 0, 0, 0)
	init_attack_counts(num_rows)
	success := place_queens_3(board, num_rows, 0, 0, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dump_board(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
