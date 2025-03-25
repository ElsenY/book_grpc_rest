package routes

import (
	"APIGateway/middleware"
	bookPb "APIGateway/pb/BookService"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type BookRoute struct {
	client bookPb.BookClient
}

func RegisterBookRoute(mainRoute *gin.Engine, bookConn *grpc.ClientConn) *gin.Engine {

	bookClient := bookPb.NewBookClient(bookConn)

	bookRoute := BookRoute{
		client: bookClient,
	}

	protectedBooks := mainRoute.Group("/books/protected")

	protectedBooks.Use(middleware.AuthMiddleware())
	protectedBooks.POST("/", bookRoute.InsertBook)
	protectedBooks.POST("/borrow-book", bookRoute.BorrowBook)
	protectedBooks.POST("/return-book", bookRoute.ReturnBook)

	return mainRoute
}

func (br BookRoute) InsertBook(c *gin.Context) {
	var reqBody bookPb.InsertBookRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, ok := c.Get("email")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no email attached"})
		return
	}

	res, err := br.client.InsertBook(c, &bookPb.InsertBookRequest{Title: reqBody.Title, Stock: reqBody.Stock, AuthorEmail: email.(string)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}

func (br BookRoute) BorrowBook(c *gin.Context) {
	var reqBody bookPb.InsertBookRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, ok := c.Get("email")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no email attached"})
		return
	}

	res, err := br.client.BorrowBook(context.Background(), &bookPb.BorrowBookRequest{Title: reqBody.Title, BorrowerEmail: email.(string)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}

func (br BookRoute) ReturnBook(c *gin.Context) {
	var reqBody bookPb.InsertBookRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, ok := c.Get("email")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no email attached"})
		return
	}

	res, err := br.client.ReturnBook(context.Background(), &bookPb.ReturnBookRequest{Title: reqBody.Title, ReturnerEmail: email.(string)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}
