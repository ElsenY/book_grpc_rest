package routes

import (
	"APIGateway/middleware"
	authorPb "APIGateway/pb/AuthorService"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	protectedAuthor.POST("/author", authorRoute.RegisterUserAsAuthor)

	return mainRoute
}

func (ar *AuthorRoute) RegisterUserAsAuthor(c *gin.Context) {

	token := c.GetHeader("Authorization")

	ctx := metadata.AppendToOutgoingContext(c, "authorization", token)

	resp, err := ar.client.RegisterUserAsAuthor(ctx, &authorPb.RegisterUserAsAuthorRequest{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}
