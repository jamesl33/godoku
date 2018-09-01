package main

import (
    "fmt"
)

// sudoku solver type which contains valid information needed to solve
// a sudoku puzzle.
type Godoku struct {
    board [][]int
    stack [][]int
}

// Although a sudoku puzzle is 9x9 it is also broken up into 3x3 logical
// this function allows the fetching of a single one of these blocks.
func (this *Godoku) GetBlock(h_pos, w_pos int) [][]int {
    var block [][]int

    for h := h_pos; h < h_pos + 3; h++ {
        block = append(block, this.board[h][w_pos:w_pos + 3])
    }

    return block
}

// Check to see if placing 'value' at 'h_pos', 'w_pos' would be valid.
func (this *Godoku) CheckPlace(h_pos, w_pos, value int) bool {
    // Check all positions in the same column
    for w := 0; w < 9; w++ {
        if this.board[h_pos][w] == value {
            return false
        }
    }

    // Check all positions in the same row
    for h := 0; h < 9; h++ {
        if this.board[h][w_pos] == value {
            return false
        }
    }

    block := this.GetBlock(int(h_pos / 3) * 3, int(w_pos / 3) * 3)

    // Check the block the placement is in
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if block[i][j] == value {
                return false
            }
        }
    }

    return true
}

// Recursively solve the sudoku puzzle.
func (this *Godoku) Solve(h_pos, w_pos, start int) bool {
    // Do not edit any places that are non-zero
    if this.board[h_pos][w_pos] != 0 {
        if w_pos < 8 {
            return this.Solve(h_pos, w_pos + 1, 1)
        } else if h_pos < 8 {
            return this.Solve(h_pos + 1, 0, 1)
        } else {
            return true
        }
    }

    // Try values until a placement is successful
    for value := start; value <= 9; value++ {
        if this.CheckPlace(h_pos, w_pos, value) {
            this.board[h_pos][w_pos] = value

            this.stack = append(this.stack, []int{h_pos, w_pos})

            if w_pos < 8 {
                return this.Solve(h_pos, w_pos + 1, 1)
            } else if h_pos < 8 {
                return this.Solve(h_pos + 1, 0, 1)
            } else {
                return true
            }
        }
    }

    // If we have reached this section it means that we have reached a point
    // where no more places could be move so we backtrack by incrementing the
    // previous placement.

    last_place := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    last_value := this.board[last_place[0]][last_place[1]]
    this.board[last_place[0]][last_place[1]] = 0

    return this.Solve(last_place[0], last_place[1], last_value + 1)
}

func main() {
    // Example taken from Wikipedia
    // Unsolved: <https://en.wikipedia.org/wiki/File:Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg>
    // Solved: <https://en.wikipedia.org/wiki/File:Sudoku_Puzzle_by_L2G-20050714_solution_standardized_layout.svg>
    godoku := Godoku{
        board: [][]int{
            {5, 3, 0, 0, 7, 0, 0, 0, 0},
            {6, 0, 0, 1, 9, 5, 0, 0, 0},
            {0, 9, 8, 0, 0, 0, 0, 6, 0},
            {8, 0, 0, 0, 6, 0, 0, 0, 3},
            {4, 0, 0, 8, 0, 3, 0, 0, 1},
            {7, 0, 0, 0, 2, 0, 0, 0, 6},
            {0, 6, 0, 0, 0, 0, 2, 8, 0},
            {0, 0, 0, 4, 1, 9, 0, 0, 5},
            {0, 0, 0, 0, 8, 0, 0, 7, 9},
        },
    }

    godoku.Solve(0, 0, 1)

    for _, row := range godoku.board {
        fmt.Println(row)
    }
}
