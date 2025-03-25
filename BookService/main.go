package main

import (
	_ "github.com/lib/pq"

	authorPb "BookService/pb/AuthorService"
	bookPb "BookService/pb/BookService"
	userPb "BookService/pb/UserService"
	"BookService/services"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbConn := InitDb()

	authorServiceConn, err := ConnectAuthorService()

	if err != nil {
		log.Println("failed to connect to user server")
	}

	defer authorServiceConn.Close()

	userServiceConn, err := ConnectUserService()

	if err != nil {
		log.Println("failed to connect to user server")
	}

	defer userServiceConn.Close()

	authorClient := authorPb.NewAuthorClient(authorServiceConn)
	userClient := userPb.NewUserClient(userServiceConn)

	bookPb.RegisterBookServer(grpcServer, &services.Server{Db: dbConn, AuthorClient: authorClient, UserClient: userClient})

	fmt.Println("🚀 Book Service is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func InitDb() *sql.DB {
	dbDsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func ConnectAuthorService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(os.Getenv("AUTHOR_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

func ConnectUserService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
