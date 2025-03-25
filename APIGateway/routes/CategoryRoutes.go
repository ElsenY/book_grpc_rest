package routes

import (
	"APIGateway/middleware"
	categoryPb "APIGateway/pb/CategoryService"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type CategoryRoute struct {
	client categoryPb.CategoryClient
}

func RegisterCategoryRoute(mainRoute *gin.Engine, categoryConn *grpc.ClientConn) *gin.Engine {

	categoryClient := categoryPb.NewCategoryClient(categoryConn)

	authorRoute := CategoryRoute{
		client: categoryClient,
	}

	protectedAuthor := mainRoute.Group("/category/protected")

	protectedAuthor.Use(middleware.AuthMiddleware())

	protectedAuthor.POST("/", authorRoute.InsertCategory)
	protectedAuthor.POST("/link-book", authorRoute.LinkBookWithCategory)

	return mainRoute
}

func (cr *CategoryRoute) InsertCategory(c *gin.Context) {
	var reqBody categoryPb.InsertCategoryRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := cr.client.InsertCategory(c, &categoryPb.InsertCategoryRequest{Name: reqBody.Name})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}

func (cr *CategoryRoute) LinkBookWithCategory(c *gin.Context) {
	var reqBody categoryPb.LinkBookWithCategoryRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := cr.client.LinkBookWithCategory(c, &categoryPb.LinkBookWithCategoryRequest{BookTitle: reqBody.BookTitle, Name: reqBody.Name})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}
