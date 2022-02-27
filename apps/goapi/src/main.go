package main

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"ndsquared/goapi/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ErrorReponse struct {
	Error string `json:"error"`
}

type ConnectFourRequest struct {
	PlayerPiece   string   `json:"player_piece" binding:"required" example:"X"`
	ComputerPiece string   `json:"computer_piece" binding:"required" example:"O"`
	EmptyPiece    string   `json:"empty_piece" binding:"required" example:"."`
	Depth         int      `json:"depth" binding:"required" example:"3"`
	Board         []string `json:"board" binding:"required" example:".....,.....,.....,.....,..X.."`
}

type ConnectfourResponse struct {
	Column int `json:"column" example:"3"`
	Value  int `json:"value" example:"5"`
}

// connectfour godoc
// @Summary      Connect Four solver
// @Description  Submit a connectfour board state and return the best move at the depth provided
// @Tags         solvers
// @Accept       json
// @Param        body  body  main.ConnectFourRequest  true  "Board State"
// @Produce      json
// @Success      200  {object}  main.ConnectfourResponse
// @Failure      400  {object}  main.ErrorReponse
// @Router       /connectfour [post]
func connectfourRoute(c *gin.Context) {
	var json ConnectFourRequest
	err := c.ShouldBindJSON(&json)
	if err != nil {
		errResponse := ErrorReponse{err.Error()}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	board := createBoard(json.Board, json.ComputerPiece, json.PlayerPiece, json.EmptyPiece)
	value, col := findBestMove(board, json.Depth, math.MinInt, math.MaxInt, json.ComputerPiece)
	r := ConnectfourResponse{Column: col, Value: value}
	c.JSON(http.StatusOK, r)
}

// @title        NDSquared GOAPI
// @version      1.0
// @description  Golang backend service for NDSquared
// @BasePath     /

func main() {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	router.Use(cors.Default())

	docs.SwaggerInfo.Host = os.Getenv("GOAPI_HOST")
	if docs.SwaggerInfo.Host == "" {
		docs.SwaggerInfo.Host = "localhost:5555"
	}
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/connectfour", connectfourRoute)

	err = router.Run("0.0.0.0:5555")
	if err != nil {
		fmt.Println(err.Error())
	}
}
