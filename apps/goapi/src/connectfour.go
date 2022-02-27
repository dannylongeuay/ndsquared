package main

import (
	"fmt"
	"strings"
)

type Board struct {
	height        int
	width         int
	cells         [][]string
	computerPiece string
	playerPiece   string
	emptyPiece    string
}

func createBoard(boardRows []string, computerPiece string, playerPiece string, emptyPiece string) Board {
	b := Board{
		computerPiece: computerPiece,
		playerPiece:   playerPiece,
		emptyPiece:    emptyPiece,
	}
	b.cells = make([][]string, len(boardRows))
	for i := range boardRows {
		b.cells[i] = strings.Split(boardRows[i], "")
	}
	b.height = len(boardRows)
	b.width = len(boardRows[0])
	return b
}

func (b *Board) getNextOpenRow(col int) (int, error) {
	for i := len(b.cells) - 1; i >= 0; i-- {
		if b.cells[i][col] == b.emptyPiece {
			return i, nil
		}
	}
	return -1, fmt.Errorf("no open row found on column %v", col)
}

func (b *Board) dropPiece(piece string, col int) error {
	row, err := b.getNextOpenRow(col)
	if err != nil {
		return err
	}
	b.cells[row][col] = piece
	return nil
}

func (b *Board) getValidMoves() []int {
	validMoves := make([]int, 0, b.width)
	for i := range b.cells[0] {
		_, err := b.getNextOpenRow(i)
		if err != nil {
			continue
		}
		validMoves = append(validMoves, i)
	}
	return validMoves
}

func (b *Board) evaluate(piece string) int {
	score := 0
	// Center columns
	centerCol := b.width / 2
	centerColWindow := b.getVerticalWindow(centerCol, 0, len(b.cells))
	centerCount := windowCount(piece, centerColWindow)
	score += centerCount * 3

	// Horizontal
	for row := range b.cells {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.cells[row][col : col+4]
			score += b.windowScore(piece, window)
		}
	}

	// Vertical
	for col := range b.cells[0] {
		for row := 0; row < len(b.cells)-3; row++ {
			window := b.getVerticalWindow(col, row, row+4)
			score += b.windowScore(piece, window)
		}
	}

	// Positive Diagonal
	for row := len(b.cells) - 1; row >= 3; row-- {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.getDiagonalWindow(col, row, -1)
			score += b.windowScore(piece, window)
		}
	}

	// Negative Diagonal
	for row := 0; row < len(b.cells)-3; row++ {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.getDiagonalWindow(col, row, 1)
			score += b.windowScore(piece, window)
		}
	}
	return score
}

func (b *Board) winningMove(piece string) bool {
	// Horizontal
	for row := range b.cells {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.cells[row][col : col+4]
			count := windowCount(piece, window)
			if count == 4 {
				return true
			}
		}
	}

	// Vertical
	for col := range b.cells[0] {
		for row := 0; row < len(b.cells)-3; row++ {
			window := b.getVerticalWindow(col, row, row+4)
			count := windowCount(piece, window)
			if count == 4 {
				return true
			}
		}
	}

	// Positive Diagonal
	for row := len(b.cells) - 1; row >= 3; row-- {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.getDiagonalWindow(col, row, -1)
			count := windowCount(piece, window)
			if count == 4 {
				return true
			}
		}
	}

	// Negative Diagonal
	for row := 0; row < len(b.cells)-3; row++ {
		for col := 0; col < len(b.cells[row])-3; col++ {
			window := b.getDiagonalWindow(col, row, 1)
			count := windowCount(piece, window)
			if count == 4 {
				return true
			}
		}
	}
	return false
}

func (b *Board) getVerticalWindow(col int, rowStart int, rowStop int) []string {
	window := make([]string, 0, 4)
	for i := rowStart; i < rowStop; i++ {
		window = append(window, b.cells[i][col])
	}
	return window
}

func (b *Board) getDiagonalWindow(colStart int, rowStart int, direction int) []string {
	window := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		window = append(window, b.cells[rowStart+i*direction][colStart+i])
	}
	return window
}

func (b *Board) copyBoard() Board {
	cb := Board{
		height:        b.height,
		width:         b.width,
		computerPiece: b.computerPiece,
		playerPiece:   b.playerPiece,
		emptyPiece:    b.emptyPiece,
	}
	cb.cells = make([][]string, len(b.cells))
	for i, row := range b.cells {
		rowCopy := make([]string, len(row))
		copy(rowCopy, row)
		cb.cells[i] = rowCopy
	}
	return cb
}

func windowCount(piece string, window []string) int {
	count := 0
	for _, p := range window {
		if piece == p {
			count++
		}
	}
	return count
}

func (b *Board) windowScore(piece string, window []string) int {
	score := 0
	playerPiece := b.computerPiece
	if piece == b.computerPiece {
		playerPiece = b.playerPiece
	}
	if windowCount(piece, window) == 4 {
		score += 100
	} else if windowCount(piece, window) == 3 && windowCount(b.emptyPiece, window) == 1 {
		score += 5
	} else if windowCount(piece, window) == 2 && windowCount(b.emptyPiece, window) == 2 {
		score += 2
	}

	if windowCount(playerPiece, window) == 3 && windowCount(b.emptyPiece, window) == 1 {
		score -= 4
	}
	return score
}
