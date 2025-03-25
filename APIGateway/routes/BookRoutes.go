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

	mainRoute.GET("/books/recommend-book", bookRoute.RecommendBook)
	mainRoute.GET("/books/search-book", bookRoute.SearchBook)
	protectedBooks := mainRoute.Group("/books/protected")

	protectedBooks.Use(middleware.AuthMiddleware())
	protectedBooks.POST("/", bookRoute.InsertBook)
	protectedBooks.POST("/borrow-book", bookRoute.BorrowBook)
	protectedBooks.POST("/return-book", bookRoute.ReturnBook)
	protectedBooks.PUT("/edit-stock", bookRoute.EditStock)

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
	var reqBody bookPb.BorrowBookRequest

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
	var reqBody bookPb.ReturnBookRequest

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

func (br BookRoute) RecommendBook(c *gin.Context) {
	var reqBody bookPb.RecommendBookRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := br.client.RecommendBook(c, &bookPb.RecommendBookRequest{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.BookTitle,
	})
}

func (br BookRoute) SearchBook(c *gin.Context) {
	var reqBody bookPb.SearchBookRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := br.client.SearchBook(c, &bookPb.SearchBookRequest{BookTitle: reqBody.BookTitle})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book_id":    res.BookId,
		"book_title": res.BookTitle,
		"stock":      res.Stock,
		"message":    res.Message,
	})
}

func (br BookRoute) EditStock(c *gin.Context) {
	var reqBody bookPb.EditBookStockRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := br.client.EditBookStock(c, &bookPb.EditBookStockRequest{Title: reqBody.Title, Stock: reqBody.Stock})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}
