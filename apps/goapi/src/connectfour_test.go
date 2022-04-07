package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoardSuite struct {
	suite.Suite
	computerPiece       string
	playerPiece         string
	emptyPiece          string
	simpleBoard         Board
	emptyBoard          Board
	fullBoard           Board
	fullColBoard        Board
	hScoreBoard         Board
	vScoreBoard         Board
	posDiagScoreBoard   Board
	negDiagScoreBoard   Board
	hWinningBoard       Board
	vWinningBoard       Board
	posDiagWinningBoard Board
	negDiagWinningBoard Board
}

func (suite *BoardSuite) SetupTest() {
	playerPiece := "O"
	suite.playerPiece = playerPiece
	computerPiece := "X"
	suite.computerPiece = computerPiece
	emptyPiece := "."
	suite.emptyPiece = emptyPiece

	simpleBoardRows := []string{".....", ".....", ".....", ".....", "..X.."}
	suite.simpleBoard = createBoard(simpleBoardRows, computerPiece, playerPiece, emptyPiece)

	emptyBoardRows := []string{".....", ".....", ".....", ".....", "....."}
	suite.emptyBoard = createBoard(emptyBoardRows, computerPiece, playerPiece, emptyPiece)

	fullBoardRows := []string{"XOXOX", "OXOXO", "XOXOX", "OXOXO", "XOXOX"}
	suite.fullBoard = createBoard(fullBoardRows, computerPiece, playerPiece, emptyPiece)

	fullColBoardRows := []string{"..X..", "..O..", "..X..", "..O..", "..X.."}
	suite.fullColBoard = createBoard(fullColBoardRows, computerPiece, playerPiece, emptyPiece)

	hScoreBoardRows := []string{".....", ".....", "XXXX.", ".....", "....."}
	suite.hScoreBoard = createBoard(hScoreBoardRows, computerPiece, playerPiece, emptyPiece)

	vScoreBoardRows := []string{".....", "..X..", "..X..", "..X..", "....."}
	suite.vScoreBoard = createBoard(vScoreBoardRows, computerPiece, playerPiece, emptyPiece)

	posDiagScoreBoardRows := []string{".....", ".....", "..X..", ".....", "X...."}
	suite.posDiagScoreBoard = createBoard(posDiagScoreBoardRows, computerPiece, playerPiece, emptyPiece)

	negDiagScoreBoardRows := []string{".....", ".O...", "..O..", "...O.", "....."}
	suite.negDiagScoreBoard = createBoard(negDiagScoreBoardRows, computerPiece, playerPiece, emptyPiece)

	hWinningBoardRows := []string{".....", ".....", "XXXX.", ".....", "....."}
	suite.hWinningBoard = createBoard(hWinningBoardRows, computerPiece, playerPiece, emptyPiece)

	vWinningBoardRows := []string{".....", "..X..", "..X..", "..X..", "..X.."}
	suite.vWinningBoard = createBoard(vWinningBoardRows, computerPiece, playerPiece, emptyPiece)

	posDiagWinningBoardRows := []string{".....", "...X.", "..X..", ".X...", "X...."}
	suite.posDiagWinningBoard = createBoard(posDiagWinningBoardRows, computerPiece, playerPiece, emptyPiece)

	negDiagWinningBoardRows := []string{".....", ".O...", "..O..", "...O.", "....O"}
	suite.negDiagWinningBoard = createBoard(negDiagWinningBoardRows, computerPiece, playerPiece, emptyPiece)
}

func (suite *BoardSuite) TestCreateBoard() {
	if suite.NotNil(suite.simpleBoard) {
		suite.Equal(5, len(suite.simpleBoard.cells))
		suite.Equal("X", suite.simpleBoard.cells[4][2])
		suite.Equal("X", suite.simpleBoard.computerPiece)
		suite.Equal("O", suite.simpleBoard.playerPiece)
		suite.Equal(".", suite.simpleBoard.emptyPiece)
	}
}

func (suite *BoardSuite) TestGetNextOpenRow() {
	row, err := suite.fullBoard.getNextOpenRow(1)
	suite.Errorf(err, "no open row found on column 1")
	suite.Equal(-1, row)

	row, err = suite.emptyBoard.getNextOpenRow(2)
	suite.Nil(err)
	suite.Equal(4, row)

	row, err = suite.simpleBoard.getNextOpenRow(2)
	suite.Nil(err)
	suite.Equal(3, row)
}

func (suite *BoardSuite) TestDropPiece() {
	err := suite.fullBoard.dropPiece(suite.computerPiece, 0)
	suite.Errorf(err, "no open row found on column 0")

	err = suite.simpleBoard.dropPiece(suite.playerPiece, 2)
	suite.Nil(err)

	suite.Equal("O", suite.simpleBoard.cells[3][2])
}

func (suite *BoardSuite) TestGetValidMoves() {
	moves := suite.fullBoard.getValidMoves()
	suite.Equal(0, len(moves))

	moves = suite.emptyBoard.getValidMoves()
	suite.EqualValues([]int{0, 1, 2, 3, 4}, moves)

	moves = suite.fullColBoard.getValidMoves()
	suite.EqualValues([]int{0, 1, 3, 4}, moves)
}

func (suite *BoardSuite) TestEvaluate() {
	score := suite.fullColBoard.evaluate(suite.computerPiece)
	suite.Equal(9, score)

	score = suite.fullColBoard.evaluate(suite.playerPiece)
	suite.Equal(6, score)

	score = suite.hScoreBoard.evaluate(suite.computerPiece)
	suite.Equal(108, score)

	score = suite.vScoreBoard.evaluate(suite.computerPiece)
	suite.Equal(19, score)

	score = suite.posDiagScoreBoard.evaluate(suite.computerPiece)
	suite.Equal(5, score)

	score = suite.negDiagScoreBoard.evaluate(suite.computerPiece)
	suite.Equal(-8, score)
}

func (suite *BoardSuite) TestWinningMove() {
	w := suite.emptyBoard.winningMove(suite.computerPiece)
	suite.False(w)

	w = suite.hWinningBoard.winningMove(suite.computerPiece)
	suite.True(w)

	w = suite.vWinningBoard.winningMove(suite.computerPiece)
	suite.True(w)

	w = suite.posDiagWinningBoard.winningMove(suite.computerPiece)
	suite.True(w)

	w = suite.negDiagWinningBoard.winningMove(suite.playerPiece)
	suite.True(w)
}

func (suite *BoardSuite) TestCopyBoard() {
	cb := suite.simpleBoard.copyBoard()
	suite.Equal(cb, suite.simpleBoard)
	suite.NotSame(cb.cells[0][0], suite.simpleBoard.cells[0][0])
	cb.cells[0][0] = "X"
	suite.NotEqual(cb, suite.simpleBoard)
	nb := cb
	suite.Same(&nb.cells[0][0], &cb.cells[0][0])
}

func TestBoardSuite(t *testing.T) {
	suite.Run(t, new(BoardSuite))
}
