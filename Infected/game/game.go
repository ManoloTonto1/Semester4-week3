package game

import (
	"encoding/json"
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

	board := Board{Size: size, Spaces: spaces}
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

func (g *Game) Start() {
	g.Turns = 0
	g.GameEnded = false
	Board := CreateEmptyBoard(5)

	for i, arr := range Board.Spaces {
		for j := range arr {
			fmt.Print(Board.Spaces[i][j])
		}
		fmt.Printf("\n")
	}

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(Board)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)
	})
}
