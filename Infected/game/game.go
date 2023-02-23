package game

import (
	"fmt"
	"log"
	"net/http"
)

type Game struct {
	Board     Board
	Turns     int
	GameEnded bool
}

func CreateEmptyBoard(size int) Board {
	spaces := make([][]int, size)
	for i := range spaces {
		spaces[i] = make([]int, size)
	}
	for i, arr := range spaces {
		for j := range arr {
			spaces[i][j] = Empty
		}
	}
	// set the players in the correct location
	spaces[size-1][0] = Player1
	spaces[0][size-1] = Player2
	players := []Player{
		{ID: Player1, Infected: 1},
		{ID: Player2, Infected: 1},
	}

	board := Board{Size: size, Spaces: spaces, Players: players}
	return board
}
func (game *Game) ServeServer() {
	port := "6969"
	//serve the server
	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) ConsoleGame(Board Board) {
	for i, arr := range Board.Spaces {
		for j := range arr {
			fmt.Print(Board.Spaces[i][j])
		}
		fmt.Printf("\n")
	}

	if Board.Players[0].Infected == 0 || Board.Players[1].Infected == 0 {
		g.GameEnded = true
	}
	if g.GameEnded {
		fmt.Println("Game ended")
		fmt.Println("winning Player infected: ", Board.Players[1].Infected)
		fmt.Println("Turns: ", g.Turns)
		fmt.Scanln()
		g.Start()
		return
	}
	fmt.Println("Player 1 turn")
	player1Move := Board.CalculateBestMoveSet(Player1, Player2)
	if err := Board.Move(Player1, player1Move); err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------")
	for i, arr := range Board.Spaces {
		for j := range arr {
			fmt.Print(Board.Spaces[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("Player 2 turn")

	aiMove := Board.CalculateBestMoveSet(Player2, Player1)

	if err := Board.Move(Player2, aiMove); err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------")
	g.Turns++
	g.ConsoleGame(Board)

}
func (g *Game) Start() {
	g.Turns = 0
	g.GameEnded = false
	Board := CreateEmptyBoard(9)

	g.ConsoleGame(Board)
}
