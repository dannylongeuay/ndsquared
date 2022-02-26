package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoardSuite struct {
	suite.Suite
	playerPiece string
	oppPiece    string
	emptyPiece  string
	simpleBoard Board
	emptyBoard  Board
	fullBoard   Board
}

func (suite *BoardSuite) SetupTest() {
	playerPiece := "X"
	suite.playerPiece = playerPiece
	oppPiece := "O"
	suite.oppPiece = oppPiece
	emptyPiece := "."
	suite.emptyPiece = emptyPiece
	simpleboard := []string{".....", ".....", ".....", ".....", "..X.."}
	suite.simpleBoard = createBoard(simpleboard, playerPiece, oppPiece, emptyPiece)

	emptyBoard := []string{".....", ".....", ".....", ".....", "....."}
	suite.emptyBoard = createBoard(emptyBoard, playerPiece, oppPiece, emptyPiece)

	fullBoard := []string{"XOXOX", "OXOXO", "XOXOX", "OXOXO", "XOXOX"}
	suite.fullBoard = createBoard(fullBoard, playerPiece, oppPiece, emptyPiece)
}

func (suite *BoardSuite) TestCreateBoard() {
	if suite.NotNil(suite.simpleBoard) {
		suite.Equal(5, len(suite.simpleBoard.cells))
		suite.Equal("X", suite.simpleBoard.cells[4][2])
		suite.Equal("X", suite.simpleBoard.playerPiece)
		suite.Equal("O", suite.simpleBoard.oppPiece)
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
	err := suite.fullBoard.dropPiece(suite.playerPiece, 0)
	suite.Errorf(err, "no open row found on column 0")

	err = suite.simpleBoard.dropPiece(suite.oppPiece, 2)
	suite.Nil(err)

	suite.Equal("O", suite.simpleBoard.cells[3][2])
}

func TestBoardSuite(t *testing.T) {
	suite.Run(t, new(BoardSuite))
}
