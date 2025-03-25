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

	// Connect to user service
	userConn, err := ConnectUserService()

	if err != nil {
		fmt.Println("failed to connect to User service", err)
	}

	defer userConn.Close()

	authorConn, err := ConnectAuthorService()

	if err != nil {
		fmt.Println("failed to connect to Author service", err)
	}

	defer userConn.Close()

	bookConn, err := ConnectBookService()

	if err != nil {
		fmt.Println("failed to connect to Book service", err)
	}

	defer bookConn.Close()

	categoryConn, err := ConnectCategoryService()

	if err != nil {
		fmt.Println("failed to connect to Category service", err)
	}

	defer categoryConn.Close()

	routes.RegisterUserRoutes(router, userConn)
	routes.RegisterAuthorRoute(router, authorConn)
	routes.RegisterBookRoute(router, bookConn)
	routes.RegisterCategoryRoute(router, categoryConn)

	fmt.Println("API Gateway running on port 8080")
	router.Run(":8080")
}

func ConnectUserService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

func ConnectAuthorService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

func ConnectBookService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

func ConnectCategoryService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
