package boards

import (
	"example/gin-api-server/ent"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Index(c *gin.Context, client *ent.Client) {
	boards, err := client.Board.Query().All(c)
	if err != nil {
		c.Error(errors.Wrap(err, ""))
	}

	c.IndentedJSON(http.StatusOK, boards)
}

func Show(c *gin.Context, client *ent.Client) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.Wrap(err, ""))
	}

	board, err := client.Board.Get(c, id)
	if err != nil {
		c.Error(errors.Wrap(err, ""))
	}

	c.IndentedJSON(http.StatusOK, board)
}

func Create(c *gin.Context, client *ent.Client) {
	var inputBoard ent.Board
	if err := c.BindJSON(&inputBoard); err != nil {
		c.Error(errors.Wrap(err, ""))
		return
	}

	newBoard, err := client.Board.Create().
		SetTitle(inputBoard.Title).
		SetContent(inputBoard.Content).
		Save(c)

	if err != nil {
		c.Error(errors.Wrap(err, ""))
		return
	}

	log.Println("board was created: ", newBoard)
	c.IndentedJSON(http.StatusOK, newBoard)
}

func Update(c *gin.Context, client *ent.Client) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.Wrap(err, ""))
	}

	var inputBoard ent.Board
	if err := c.BindJSON(&inputBoard); err != nil {
		c.Error(errors.Wrap(err, ""))
		return
	}

	updatedBoard, err := client.Board.
		UpdateOneID(id).
		SetTitle(inputBoard.Title).
		SetContent(inputBoard.Content).
		Save(c)
	if err != nil {
		c.Error(errors.Wrap(err, ""))
	}

	c.IndentedJSON(http.StatusOK, updatedBoard)
}

// curl http://localhost:8888/api/boards \
//     --include \
//     --header "Content-Type: application/json" \
//     --request "POST" \
//     --data '{"title": "첫 게시글","content": "안녕하세요!"}'

// curl http://localhost:8888/api/boards \
//     --include \
//     --header "Content-Type: application/json" \
//     --request "GET"

// curl http://localhost:8888/api/boards/3 \
//     --include \
//     --header "Content-Type: application/json" \
//     --request "GET"

// curl http://localhost:8888/api/boards/3 \
//     --include \
//     --header "Content-Type: application/json" \
//     --request "PATCH" \
//     --data '{"title": "세번째 게시글","content": "반가워요!"}'
