package main

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConnectFourRequest struct {
	PlayerPiece   string     `json:"player_piece" binding:"required"`
	ComputerPiece string     `json:"computer_piece" binding:"required"`
	EmptyPiece    string     `json:"empty_piece" binding:"required"`
	Depth         int        `json:"depth" binding:"required"`
	Board         [][]string `json:"board" binding:"required"`
}

type ConnectfourResponse struct {
	Column int `json:"column"`
	Value  int `json:"value"`
}

func connectfour(c *gin.Context) {
	var json ConnectFourRequest
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	board := createBoard(json.Board, json.PlayerPiece, json.ComputerPiece, json.EmptyPiece)
	value, col := minimax(board, json.Depth, math.MinInt, math.MaxInt, json.ComputerPiece)
	r := ConnectfourResponse{Column: col, Value: value}
	c.JSON(http.StatusOK, r)
}

func main() {
	router := gin.Default()
	router.POST("/connectfour", connectfour)
	router.Run()
}
