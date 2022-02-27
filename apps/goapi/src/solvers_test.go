package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SolversSuite struct {
	suite.Suite
	playerPiece       string
	computerPiece     string
	emptyPiece        string
	simpleBoard       Board
	emptyBoard        Board
	fullBoard         Board
	winningMoveBoard  Board
	blockingMoveBoard Board
}

func (suite *SolversSuite) SetupTest() {
	playerPiece := "X"
	suite.playerPiece = playerPiece
	computerPiece := "O"
	suite.computerPiece = computerPiece
	emptyPiece := "."
	suite.emptyPiece = emptyPiece

	simpleBoardRows := []string{".....", ".....", ".....", ".....", "..X.."}
	suite.simpleBoard = createBoard(simpleBoardRows, computerPiece, playerPiece, emptyPiece)

	emptyBoardRows := []string{".....", ".....", ".....", ".....", "....."}
	suite.emptyBoard = createBoard(emptyBoardRows, computerPiece, playerPiece, emptyPiece)

	fullBoardRows := []string{"XOXOX", "OXOXO", "XOXOX", "OXOXO", "XOXOX"}
	suite.fullBoard = createBoard(fullBoardRows, computerPiece, playerPiece, emptyPiece)

	winningMoveBoardRows := []string{".....", ".....", "..XO.", "..XO.", "X.XO."}
	suite.winningMoveBoard = createBoard(winningMoveBoardRows, computerPiece, playerPiece, emptyPiece)

	blockingMoveBoardRows := []string{".....", ".....", ".X...", ".X.O.", ".X.O."}
	suite.blockingMoveBoard = createBoard(blockingMoveBoardRows, computerPiece, playerPiece, emptyPiece)
}

func (suite *SolversSuite) TestFindBestMove() {
	val, col := findBestMove(suite.fullBoard, 3, math.MinInt, math.MaxInt, suite.computerPiece)
	suite.Equal(0, val)
	suite.Equal(-1, col)

	val, col = findBestMove(suite.emptyBoard, 3, math.MinInt, math.MaxInt, suite.computerPiece)
	suite.Equal(7, val)
	suite.Equal(2, col)

	val, col = findBestMove(suite.simpleBoard, 3, math.MinInt, math.MaxInt, suite.computerPiece)
	suite.Equal(6, val)
	suite.Equal(2, col)

	val, col = findBestMove(suite.winningMoveBoard, 3, math.MinInt, math.MaxInt, suite.computerPiece)
	suite.Equal(math.MaxInt, val)
	suite.Equal(3, col)

	val, col = findBestMove(suite.blockingMoveBoard, 3, math.MinInt, math.MaxInt, suite.computerPiece)
	suite.Equal(8, val)
	suite.Equal(1, col)
}

func TestSolversSuite(t *testing.T) {
	suite.Run(t, new(SolversSuite))
}
