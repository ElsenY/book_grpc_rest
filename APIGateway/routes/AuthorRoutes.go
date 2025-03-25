package routes

import (
	"APIGateway/middleware"
	authorPb "APIGateway/pb/AuthorService"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AuthorRoute struct {
	client authorPb.AuthorClient
}

func RegisterAuthorRoute(mainRoute *gin.Engine, authorConn *grpc.ClientConn) *gin.Engine {

	authorClient := authorPb.NewAuthorClient(authorConn)

	authorRoute := AuthorRoute{
		client: authorClient,
	}

	protectedAuthor := mainRoute.Group("/authors/protected")

	protectedAuthor.Use(middleware.AuthMiddleware())

	protectedAuthor.POST("/", authorRoute.RegisterUserAsAuthor)

	return mainRoute
}

func (ar *AuthorRoute) RegisterUserAsAuthor(c *gin.Context) {

	email, ok := c.Get("email")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no email attached"})
		return
	}

	resp, err := ar.client.RegisterUserAsAuthor(c, &authorPb.RegisterUserAsAuthorRequest{UserEmail: email.(string)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}
