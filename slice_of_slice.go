package main

import (
	"fmt"
	"strings"
)

func main() {
	// create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	game[0][2] = "X"
	game[1][0] = "O"
	game[1][1] = "X"
	game[2][0] = "X"
	game[2][2] = "O"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
