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
	spaces[size-1][1] = Player1
	spaces[size-2][0] = Player1
	spaces[size-2][1] = Player1

	spaces[0][size-2] = Player2
	spaces[1][size-2] = Player2
	spaces[0][size-1] = Player2
	spaces[1][size-1] = Player2
	players := []Player{
		{ID: Player1, Infected: 4},
		{ID: Player2, Infected: 4},
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
	Board.PrintBoard()
	CheckGameFinished(Board, g)

	fmt.Println("Player 1 turn")
	player1Move := Board.CalculateBestMoveSet(Player1, Player2)
	if err := Board.Move(Player1, player1Move); err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------")
	Board.PrintBoard()
	CheckGameFinished(Board, g)

	fmt.Println("Player 2 turn")

	aiMove := Board.CalculateBestMoveSet(Player2, Player1)

	if err := Board.Move(Player2, aiMove); err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------")
	g.Turns++
	g.ConsoleGame(Board)

}

func CheckGameFinished(Board Board, g *Game) {
	if Board.Players[0].Infected == 0 || Board.Players[1].Infected == 0 {
		g.GameEnded = true
	}
	if g.GameEnded {
		fmt.Println("Game ended")
		fmt.Println("winning Player infected: ", Board.Players[1].Infected)
		fmt.Println("Turns: ", g.Turns)
		fmt.Scanln()
		g.Start()
	}
}

func (Board Board) PrintBoard() {
	print("-|")
	for i := 0; i < Board.Size; i++ {
		fmt.Print(i)
	}
	fmt.Printf("\n")
	for i := 0; i < Board.Size+2; i++ {

		fmt.Print("-")
	}
	fmt.Printf("\n")
	for i, arr := range Board.Spaces {
		fmt.Print(i)
		fmt.Print("|")
		for j := range arr {
			fmt.Print(Board.Spaces[i][j])
		}
		fmt.Printf("\n")
	}
}
func (g *Game) PlayerGame(Board Board) {

	Board.PrintBoard()
	CheckGameFinished(Board, g)

	fmt.Println("Player 1 turn")
	fmt.Println("Enter the coordinates of the piece you want to move (y,x)")
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println("Enter the coordinates of the space you want to move to (y,x)")
	var x2, y2 int
	fmt.Scanln(&x2, &y2)
	player1Move := Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x2, Y: y2}}
	if err := Board.Move(Player1, player1Move); err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------")
	Board.PrintBoard()
	CheckGameFinished(Board, g)

	fmt.Println("Player 2 turn")

	aiMove := Board.CalculateBestMoveSet(Player2, Player1)

	if err := Board.Move(Player2, aiMove); err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------")
	g.Turns++
	g.PlayerGame(Board)
}

func (g *Game) Start() {
	g.Turns = 0
	g.GameEnded = false
	Board := CreateEmptyBoard(9)
	fmt.Println("------------------")
	fmt.Println("Choose a game mode:")
	fmt.Println("1. Console")
	fmt.Println("2. P vs M")
	fmt.Println("------------------")
	var gameMode int
	fmt.Scanln(&gameMode)
	if gameMode == 1 {
		g.ConsoleGame(Board)
	}
	if gameMode == 2 {
		g.PlayerGame(Board)
	}
}
