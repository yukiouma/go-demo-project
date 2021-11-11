package book

import (
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BookHttpServer interface {
	FindOneBook(q *QueryBookReq) (*BookRes, error)
	SaleOneBook(q *QueryBookReq) (*BookRes, error)
	NewOneBook(q *SaveBookReq) (*BookRes, error)
	DeleteOneBook(q *QueryBookReq) (*DeleteBookRes, error)
}

func RegisterBookHttpServer(engine *gin.Engine, server BookHttpServer) {
	g := engine.Group(prefix())
	{
		g.GET("/find/:id", FindOneBookTransfer(server.FindOneBook))
		g.POST("/new", NewOneBookTransfer(server.NewOneBook))
		g.GET("/sale/:id", SaleOneBookTransfer(server.SaleOneBook))
		g.DELETE("/delete/:id", DeleteOneBookTransfer(server.DeleteOneBook))
	}
}

// get api prefix according file location
func prefix() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to assign router prefix")
	}
	dirFrag := strings.Split(strings.Replace(file, dir, "", -1), "/")
	dirFrag = dirFrag[:len(dirFrag)-1]
	return strings.Join(dirFrag, "/")
}

// request and response parameter definitions
type QueryBookReq struct {
	ID int
}

type SaveBookReq struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"reuqired"`
	Saled bool   `json:"saled" binding:"reuqired"`
}

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Saled bool   `json:"saled"`
}

type BookRes struct {
	Data    *Book  `json:"data"`
	Message string `json:"message"`
}

type DeleteBookRes struct {
	Message string `json:"message"`
}

// transfer restful request method to gin handler functions
func FindOneBookTransfer(f func(q *QueryBookReq) (*BookRes, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := new(QueryBookReq)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		q.ID = id
		book, err := f(q)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func SaleOneBookTransfer(f func(q *QueryBookReq) (*BookRes, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := new(QueryBookReq)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		q.ID = id
		book, err := f(q)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func DeleteOneBookTransfer(f func(q *QueryBookReq) (*DeleteBookRes, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := new(QueryBookReq)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		q.ID = id
		book, err := f(q)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func NewOneBookTransfer(f func(q *SaveBookReq) (*BookRes, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q SaveBookReq
		if err := c.ShouldBind(&q); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		book, err := f(&q)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}
