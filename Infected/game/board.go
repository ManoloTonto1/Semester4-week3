package game

import (
	"errors"
	"sync"
)

type Board struct {
	Players []Player
	Size    int
	Spaces  [][]int `json:"spaces"`
}
type Player struct {
	ID       PlayerID
	Infected int
}
type PlayerID = int

const (
	Empty   PlayerID = 0
	Player1 PlayerID = 1
	Player2 PlayerID = 2
)

type Move struct {
	From          Coordinates `json:"from"`
	To            Coordinates `json:"to"`
	Weight        int
	boardInstance [][]int
}
type Coordinates struct {
	X int
	Y int
}

func (b *Board) CheckMovesPossiblePerInstance(x, y int) []Move {
	var wg sync.WaitGroup
	var moves []Move

	// Infection
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				if x == i && y == j {
					return
				}
				if x+i > b.Size-1 || x+i < 0 {
					return
				}
				if y+j > b.Size-1 || y+j < 0 {
					return
				}
				if b.Spaces[x+i][y+j] == Empty {
					moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + i, Y: y + j}})
				}
			}(i, j)
		}
	}
	wg.Wait()

	return moves
}
func (b *Board) CountInstances(player PlayerID) int {
	count := 0
	for x, array := range b.Spaces {
		for y := range array {
			if b.Spaces[x][y] == player {
				count++
			}
		}
	}
	return count
}
func (b *Board) GetPossibleMoves(playerType PlayerID) []Move {
	var wg sync.WaitGroup
	var possibleMoves []Move

	for x, array := range b.Spaces {
		for y := range array {
			wg.Add(1)
			go func(x int, y int) {
				defer wg.Done()
				if b.Spaces[x][y] == playerType {
					calcMoves := b.CheckMovesPossiblePerInstance(x, y)
					possibleMoves = append(possibleMoves, calcMoves...)
				}
			}(x, y)
		}
	}
	wg.Wait()
	return possibleMoves
}

func (b *Board) Infect(move Move, player PlayerID, enemy PlayerID) (int, [][]int) {
	var wg sync.WaitGroup
	infectionCount := 0
	newBoardInstance := make([][]int, len(b.Spaces))
	copy(newBoardInstance, b.Spaces)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				if move.To.X == move.From.X && move.To.Y == move.From.Y {
					return
				}
				if move.To.X+i > b.Size-1 || move.From.X < 0 || move.To.X+i < 0 {
					return
				}
				if move.To.Y+j > b.Size-1 || move.From.Y < 0 || move.To.Y+j < 0 {
					return
				}
				if b.Spaces[move.To.X+i][move.To.Y+j] == enemy {
					infectionCount++
					newBoardInstance[move.To.X+i][move.To.Y+j] = player
					return
				}
			}(i, j)
		}
	}
	wg.Wait()
	return infectionCount, newBoardInstance
}

func (b *Board) CalculateWeightsPerMove(player PlayerID, enemy PlayerID) (Move, []Move) {
	var wg sync.WaitGroup
	moves := b.GetPossibleMoves(player)
	var bestMove Move

	for index, move := range moves {
		wg.Add(1)
		go func(index int, move Move) {
			defer wg.Done()
			weight, newBoardInstance := b.Infect(move, player, enemy)
			moves[index].Weight = weight
			moves[index].boardInstance = newBoardInstance

			if bestMove.Weight == 0 {
				bestMove.Weight = weight
			}
			if weight > bestMove.Weight {
				bestMove.Weight = weight
			}
		}(index, move)
	}
	wg.Wait()
	return bestMove, moves
}
func (b *Board) CalculateBestMoveSet(player, enemy PlayerID) Move {
	var wg sync.WaitGroup

	_, FirstMoves := b.CalculateWeightsPerMove(player, enemy)
	var bestMovePossible Move
	for _, move := range FirstMoves {
		wg.Add(1)
		go func(move Move) {
			defer wg.Done()
			virtualBoard := Board{Spaces: move.boardInstance}
			bestMove, _ := virtualBoard.CalculateWeightsPerMove(enemy, player)
			move.Weight -= bestMove.Weight
			if bestMovePossible.Weight == 0 {
				bestMovePossible = move
			}
			if move.Weight > bestMovePossible.Weight {
				bestMovePossible = move
			}
		}(move)

	}
	wg.Wait()

	return bestMovePossible
}

func DetermineEnemyPlayer(player PlayerID) PlayerID {
	enemyPlayer := Player2

	if player == enemyPlayer {
		enemyPlayer = Player1
	}
	return enemyPlayer
}
func (b *Board) UpdateInfections(playerID int) {
	var wg sync.WaitGroup
	b.Players[0].Infected = 0
	b.Players[1].Infected = 0
	for index, arr := range b.Spaces {
		for index2 := range arr {
			wg.Add(1)
			go func(index int, index2 int) {
				defer wg.Done()
				if b.Spaces[index][index2] == playerID {
					b.Players[0].Infected++
				}

				if b.Spaces[index][index2] == DetermineEnemyPlayer(playerID) {
					b.Players[1].Infected++
				}
			}(index, index2)
		}
	}
	wg.Wait()
}
func (b *Board) Move(player PlayerID, move Move) error {
	enemyPlayer := DetermineEnemyPlayer(player)
	if b.Spaces[move.From.X][move.From.Y] != player {
		return errors.New("cannot move from empty space")
	}

	if b.Spaces[move.To.X][move.To.Y] == enemyPlayer {
		return errors.New("cannot move to enemy space")
	}
	if b.Spaces[move.To.X][move.To.Y] != Empty {
		return errors.New("space is not empty")
	}
	b.Spaces[move.To.X][move.To.Y] = player

	dy := move.To.Y - move.From.Y
	dx := move.To.X - move.From.X
	if dx == -2 || dx == 2 ||
		dy == -2 || dy == 2 {
		b.Spaces[move.From.X][move.From.Y] = Empty
	}

	b.Infect(move, player, enemyPlayer)
	b.UpdateInfections(player)
	b.UpdateInfections(enemyPlayer)
	return nil
}
