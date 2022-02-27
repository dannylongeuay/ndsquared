package main

import (
	"fmt"
	"math"
	"math/rand"
)

// func (b *Board) pickBestMove(piece string) (int, error) {
// 	bestScore := math.MinInt
// 	bestCol := -1
// 	for i := range b.cells[0] {
// 		tmpBoard := b.copyBoard()
// 		err := tmpBoard.dropPiece(piece, i)
// 		if err != nil {
// 			fmt.Println("tried to drop piece in invalid col")
// 			continue
// 		}
// 		score := tmpBoard.evaluate(piece)
// 		if score > bestScore {
// 			bestScore = score
// 			bestCol = i
// 		}
// 	}
// 	if bestCol == -1 {
// 		return -1, fmt.Errorf("no valid move could be found for piece %v", piece)
// 	}
// 	return bestCol, nil
// }

func minimax(board Board, depth int, alpha int, beta int, maximizingComputer string) (int, int) {
	returnValue := math.MaxInt
	if maximizingComputer == board.computerPiece {
		returnValue = math.MinInt
	}
	computerWon := board.winningMove(board.computerPiece)
	oppWon := board.winningMove(board.oppPiece)
	validMoves := board.getValidMoves()
	if depth == 0 || len(validMoves) == 0 || computerWon || oppWon {
		if computerWon {
			return math.MaxInt, -1
		} else if oppWon {
			return math.MinInt, -1
		} else if depth == 0 {
			return board.evaluate(board.computerPiece), -1
		} else {
			return 0, -1
		}
	}
	returnCol := validMoves[rand.Intn(len(validMoves))]
	for _, col := range validMoves {
		newBoard := board.copyBoard()
		err := newBoard.dropPiece(maximizingComputer, col)
		if err != nil {
			fmt.Println(err.Error())
			return -1, -1
		}
		nextComputer := board.oppPiece
		if maximizingComputer == board.oppPiece {
			nextComputer = board.computerPiece
		}
		value, _ := minimax(newBoard, depth-1, alpha, beta, nextComputer)
		if maximizingComputer == board.computerPiece {
			if value > returnValue {
				returnValue = value
				returnCol = col
			}
			if returnValue > alpha {
				alpha = returnValue
			}
			if alpha >= beta {
				break
			}
		} else {
			if value < returnValue {
				returnValue = value
				returnCol = col
			}
			if returnValue < beta {
				beta = value
			}
			if beta <= alpha {
				break
			}
		}
	}
	return returnValue, returnCol
}
