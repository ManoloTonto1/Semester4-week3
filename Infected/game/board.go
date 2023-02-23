package game

import "sync"

type Board struct {
	Size   int
	Spaces [][]int `json:"spaces"`
}
type Player = int

const (
	Empty   Player = 0
	Player1 Player = 1
	Player2 Player = 2
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

	wg.Add(16)
	// Infection
	go func() {
		defer wg.Done()
		if x != 0 && y != 0 && b.Spaces[x-1][y-1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 1, Y: y - 1}})
		}
	}()
	go func() {
		defer wg.Done()
		if x != 0 && b.Spaces[x-1][y] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 1, Y: y}})
		}
	}()
	go func() {
		defer wg.Done()
		if y != 0 && b.Spaces[x][y-1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x, Y: y - 1}})
		}
	}()
	go func() {
		defer wg.Done()
		if x < b.Size-1 && y > b.Size-1 && b.Spaces[x+1][y+1] == Empty {

			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 1, Y: y + 1}})
		}
	}()
	go func() {
		defer wg.Done()

		if x < b.Size-1 && b.Spaces[x+1][y] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 1, Y: y}})
		}
	}()
	go func() {
		defer wg.Done()

		if y < b.Size-1 && b.Spaces[x][y+1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x, Y: y + 1}})
		}
	}()
	go func() {
		defer wg.Done()

		if x != 0 && y < b.Size-1 && b.Spaces[x-1][y+1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 1, Y: y + 1}})
		}
	}()
	go func() {
		defer wg.Done()

		if x < b.Size-1 && y != 0 && b.Spaces[x+1][y-1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 1, Y: y - 1}})
		}
	}()

	// Straight Jumps
	go func() {
		defer wg.Done()
		if x > 1 && b.Spaces[x-2][y] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 2, Y: y}})
		}
	}()
	go func() {
		defer wg.Done()
		if y > 1 && b.Spaces[x][y-2] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x, Y: y - 2}})
		}
	}()
	go func() {
		defer wg.Done()
		if x < b.Size-2 && b.Spaces[x+2][y] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 2, Y: y}})
		}
	}()
	go func() {
		defer wg.Done()
		if y < b.Size-2 && b.Spaces[x][y+2] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x, Y: y + 2}})
		}
	}()

	// Diagonal Jumps
	go func() {
		defer wg.Done()
		if x > 1 && y > 0 && b.Spaces[x-2][y-1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 2, Y: y - 1}})
		}
	}()
	go func() {
		defer wg.Done()
		if y > 1 && x < b.Size-1 && b.Spaces[x-2][y+1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x - 2, Y: y + 1}})
		}
	}()
	go func() {
		defer wg.Done()
		if x < b.Size-2 && y > 0 && b.Spaces[x+2][y-1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 2, Y: y - 1}})
		}
	}()
	go func() {
		defer wg.Done()
		if y < b.Size-2 && x > b.Size-1 && b.Spaces[x+2][y+1] == Empty {
			moves = append(moves, Move{From: Coordinates{X: x, Y: y}, To: Coordinates{X: x + 2, Y: y + 1}})
		}
	}()
	wg.Wait()

	return moves
}
func (b *Board) CountPossibleMoves(player Player) int {
	count := 0
	for x, array := range b.Spaces {
		for y := range array {
			if b.Spaces[x][y] == player {
				count += 2

			}
		}
	}
	return count
}
func (b *Board) GetPossibleMoves(playerType Player) []Move {
	var wg sync.WaitGroup
	var possibleMoves []Move

	wg.Add(b.CountPossibleMoves(playerType))
	for x, array := range b.Spaces {
		for y := range array {
			go func(_x int, _y int) {
				defer wg.Done()
				if b.Spaces[_x][_y] == playerType {
					calcMoves := b.CheckMovesPossiblePerInstance(_x, _y)
					possibleMoves = append(possibleMoves, calcMoves...)
				}
			}(x, y)
		}
	}
	wg.Wait()
	return possibleMoves
}

func (b *Board) Infect(move Move, player Player, enemy Player) (int, [][]int) {
	var wg sync.WaitGroup
	infectionCount := 0
	newBoardInstance := make([][]int, len(b.Spaces))
	copy(newBoardInstance, b.Spaces)
	wg.Add(9)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			go func(_i int, _j int) {
				defer wg.Done()
				if move.To.X == move.From.X && move.To.Y == move.From.Y {
					return
				}
				if move.To.X+_i > b.Size || move.From.X < 0 {
					return
				}
				if move.To.Y+_j > b.Size || move.From.Y < 0 {
					return
				}
				if b.Spaces[move.To.X+_i][move.To.Y+_j] == enemy {
					infectionCount++
					newBoardInstance[move.To.X+_i][move.To.Y+_j] = player
				}
			}(i, j)
		}
	}
	wg.Wait()
	return infectionCount, newBoardInstance
}

func (b *Board) CalculateWeightsPerMove(player Player, enemy Player) (Move, []Move) {
	var wg sync.WaitGroup
	moves := b.GetPossibleMoves(player)
	var bestMove Move
	wg.Add(len(moves))

	for index, move := range moves {
		go func(_index int, _move Move) {
			defer wg.Done()
			weight, newBoardInstance := b.Infect(_move, player, enemy)
			moves[_index].Weight = weight
			moves[_index].boardInstance = newBoardInstance

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
func (b *Board) CalculateBestMoveSet(player, enemy Player) Move {
	var wg sync.WaitGroup

	_, FirstMoves := b.CalculateWeightsPerMove(player, enemy)
	var bestMovePossible Move
	wg.Add(len(FirstMoves))
	for _, move := range FirstMoves {
		go func(_move *Move) {
			defer wg.Done()
			virtualBoard := Board{Spaces: _move.boardInstance}
			bestMove, _ := virtualBoard.CalculateWeightsPerMove(enemy, player)
			_move.Weight -= bestMove.Weight
			if _move.Weight > bestMovePossible.Weight {
				bestMovePossible = *_move
			}
		}(&move)

	}
	wg.Wait()

	return bestMovePossible
}

func DetermineEnemyPlayer(player Player) Player {
	enemyPlayer := Player2

	if player == enemyPlayer {
		enemyPlayer = Player1
	}
	return enemyPlayer
}

func (b *Board) Move(player Player, move Move) {
	if b.Spaces[move.To.X][move.To.Y] != Empty {
		return
	}
	b.Spaces[move.To.X][move.To.Y] = player

	if move.To.X == -2 || move.To.X == 2 ||
		move.To.Y == -2 || move.To.Y == 2 {
		b.Spaces[move.From.X][move.From.Y] = Empty
	}

	b.Infect(move, player, DetermineEnemyPlayer(player))
}
