package routes

import (
	userPb "APIGateway/UserService/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserRoute struct {
	client userPb.UserClient
}

func RegisterUserRoutes(mainRoute *gin.Engine, userConn *grpc.ClientConn) *gin.Engine {

	userClient := userPb.NewUserClient(userConn)

	userRoute := UserRoute{
		client: userClient,
	}

	mainRoute.GET("/users/:id", userRoute.GetUser)
	mainRoute.POST("/users", userRoute.RegisterUser)
	mainRoute.POST("/login", userRoute.Login)

	return mainRoute
}

func (ur UserRoute) GetUser(c *gin.Context) {
	id := c.Param("id")

	res, err := ur.client.GetUser(context.Background(), &userPb.GetUserRequest{UserId: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}

func (ur UserRoute) RegisterUser(c *gin.Context) {
	var reqBody userPb.RegisterUserRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqBody.Password == "" || reqBody.Name == "" || reqBody.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	res, err := ur.client.RegisterUser(context.Background(), &userPb.RegisterUserRequest{Name: reqBody.Name, Password: reqBody.Password, Email: reqBody.Email})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}

func (ur UserRoute) Login(c *gin.Context) {
	var reqBody userPb.LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ur.client.Login(context.Background(), &userPb.LoginRequest{Password: reqBody.Password, Email: reqBody.Email})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": res.Token,
	})
}
