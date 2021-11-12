package v1

import (
	context "context"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterBookHttpServer(engine *gin.Engine, server BookServiceServer) {
	g := engine.Group(Prefix())
	// g := engine.Group("/api/book/v1/")
	{
		g.GET("/find/:id", FindBookTransfer(server.FindBook))
		g.POST("/new", NewBookTransfer(server.NewBook))
		g.GET("/sale/:id", SaleBookTransfer(server.SaleBook))
		g.DELETE("/delete/:id", DeleteBookTransfer(server.DeleteBook))
	}
}

// get api prefix according file location
func Prefix() string {
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

// transfer restful request method to gin handler functions
func FindBookTransfer(f func(ctx context.Context, in *FindBookRequest) (*BookReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(FindBookRequest)
		// id, err := strconv.Atoi(c.Param("id"))
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		book, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func SaleBookTransfer(f func(ctx context.Context, in *SaleBookRequest) (*BookReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(SaleBookRequest)
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		book, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func DeleteBookTransfer(f func(ctx context.Context, in *DeleteBookRequest) (*DeleteBookReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(DeleteBookRequest)
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		book, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}

func NewBookTransfer(f func(ctx context.Context, in *NewBookRequest) (*BookReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in NewBookRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		book, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, book)
	}
}
