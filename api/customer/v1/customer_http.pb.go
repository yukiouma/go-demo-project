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

func RegisterCustomerHttpServer(engine *gin.Engine, server CustomerServiceServer) {
	g := engine.Group(Prefix())
	// g := engine.Group("/api/book/v1/")
	{
		g.GET("/find/:id", FindCustomerTransfer(server.FindCustomer))
		g.POST("/register", RegisterCustomerTransfer(server.RegisterCustomer))
		g.POST("/update", UpdateCustomerTransfer(server.UpdateCustomer))
		g.DELETE("/remove/:id", RemoveCustomerTransfer(server.RemoveCustomer))
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

func FindCustomerTransfer(f func(ctx context.Context, in *FindCustomerRequest) (*CustomerReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(FindCustomerRequest)
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		customer, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, customer)
	}
}

func RegisterCustomerTransfer(f func(ctx context.Context, in *RegisterCustomerRequest) (*CustomerReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in RegisterCustomerRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		customer, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, customer)
	}
}

func UpdateCustomerTransfer(f func(ctx context.Context, in *Customer) (*CustomerReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in Customer
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		customer, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, customer)
	}
}

func RemoveCustomerTransfer(f func(ctx context.Context, in *FindCustomerRequest) (*RemoveCustomerReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(FindCustomerRequest)
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		customer, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, customer)
	}
}
