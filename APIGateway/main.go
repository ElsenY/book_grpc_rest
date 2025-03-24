package main

import (
	"APIGateway/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	router := gin.Default()

	// Connect to gRPC service
	userConn, err := ConnectUserService()

	if err != nil {
		fmt.Println("failed to connect to user service")
	}

	routes.RegisterUserRoutes(router, userConn)

	fmt.Println("API Gateway running on port 8080")
	router.Run(":8080")
}

func ConnectUserService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
