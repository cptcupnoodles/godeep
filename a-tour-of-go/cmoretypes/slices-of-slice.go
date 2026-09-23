package cmoretypes

import (
	"fmt"
)

func TicTacToe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i, v := range board {
		fmt.Printf("%s\n", i, v, " ")
		//	fmt.Printf("2**%d = %d\n", i, v)
	}
}
