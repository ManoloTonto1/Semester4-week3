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
	fmt.Println("------------------")
	fmt.Println("Player 1 infected: ", Board.Players[0].Infected)
	fmt.Println("Player 2 infected: ", Board.Players[1].Infected)
	fmt.Println("Turns: ", g.Turns)
	fmt.Println("------------------")
	fmt.Scanln()

	if Board.Players[0].Infected == 0 || Board.Players[1].Infected == 0 {
		g.GameEnded = true
	}
	if g.GameEnded {
		fmt.Println("Game ended")
		fmt.Println("Player 1 infected: ", Board.Players[0].Infected)
		fmt.Println("Player 2 infected: ", Board.Players[1].Infected)
		fmt.Println("Turns: ", g.Turns)
		fmt.Scanln()
		g.Start()
		return
	}

	fmt.Println("Player 1 turn")
	// fmt.Println("Choose a block to move (x y):")
	// // choose a block to move
	// var x, y int
	// fmt.Scan(&x, &y)
	// fmt.Println("Choose a direction to move (x y):")
	// // choose a direction to move
	// var toX, toY int
	// fmt.Scan(&toX, &toY)

	// err := Board.Move(Player1, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: toX, Y: toY}})
	player1Move := Board.CalculateBestMoveSet(Player1, Player2)
	if err := Board.Move(Player1, player1Move); err != nil {
		fmt.Println(err)
		g.ConsoleGame(Board)
		return
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
		g.ConsoleGame(Board)
		return
	}
	fmt.Println("------------------")
	g.Turns++
	g.ConsoleGame(Board)

}
func (g *Game) Start() {
	g.Turns = 0
	g.GameEnded = false
	Board := CreateEmptyBoard(7)

	g.ConsoleGame(Board)

	// http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
	// 	b, err := json.Marshal(Board)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	w.Write(b)
	// })

	// http.HandleFunc("/game/move", func(w http.ResponseWriter, r *http.Request) {
	// 	// Player turn
	// 	var move Move
	// 	err := json.NewDecoder(r.Body).Decode(&move)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	Board.Move(Player1, move)
	// 	b, err := json.Marshal(Board)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	// AI turn
	// 	aiMove := Board.CalculateBestMoveSet(Player2, Player1)
	// 	Board.Move(Player2, aiMove)

	// 	w.Write(b)
	// })
	// g.ServeServer()

}
