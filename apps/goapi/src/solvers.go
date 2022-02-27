package main

import (
	"fmt"
	"math"
	"math/rand"
)

func findBestMove(board Board, depth int, alpha int, beta int, maximizingPiece string) (int, int) {
	validMoves := board.getValidMoves()

	if len(validMoves) == 0 {
		return 0, -1
	}

	for _, col := range validMoves {
		newBoard := board.copyBoard()
		err := newBoard.dropPiece(maximizingPiece, col)
		if err != nil {
			continue
		}
		if newBoard.winningMove(maximizingPiece) {
			return math.MaxInt, col
		}
	}

	return minimax(board, depth, alpha, beta, maximizingPiece)
}

func minimax(board Board, depth int, alpha int, beta int, maximizingPiece string) (int, int) {
	returnValue := math.MaxInt
	if maximizingPiece == board.computerPiece {
		returnValue = math.MinInt
	}
	computerWon := board.winningMove(board.computerPiece)
	playerWon := board.winningMove(board.playerPiece)
	validMoves := board.getValidMoves()
	if depth == 0 || len(validMoves) == 0 || computerWon || playerWon {
		if computerWon {
			return math.MaxInt, -1
		} else if playerWon {
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
		err := newBoard.dropPiece(maximizingPiece, col)
		if err != nil {
			fmt.Println(err.Error())
			return -1, -1
		}
		if newBoard.winningMove(maximizingPiece) {
			returnValue = math.MaxInt
			returnCol = col
			if maximizingPiece == board.playerPiece {
				returnValue = math.MinInt
			}
			break
		}
		nextPiece := board.playerPiece
		if maximizingPiece == board.playerPiece {
			nextPiece = board.computerPiece
		}
		value, _ := minimax(newBoard, depth-1, alpha, beta, nextPiece)
		if maximizingPiece == board.computerPiece {
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
