package main

import (
	"fmt"
	"mine-hunter/mines"
)

func main() {
	fmt.Println("\033[2J")
	game := mines.NewGame(20, 30, 20)

	for !game.Ended {
		game.Print()

		game.Input()
	}

	game.Print()
	game.GoodBye()
}
